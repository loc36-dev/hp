package server

import (
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.3.0
	"gopkg.in/qamarian-dtp/squaket.v0" // v0.1.1
	"strings"
)

type requestData struct {}

// validate () checks if the request data of the client is valid. If the request data is not valid, the client's request would not be served.
func (d *requestData) validate () (*validatedRequestData) {
	// Retrieval of request data. ...2... {
	data, _ := mux.Vars ()["locations"]
	// ...2... }

	// Checking if request data was properly formatted. ...2... {
	if data == "" {
		panic (invErr0)
	} else if len (data) > 1024 {
		panic (invErr5)
	}

	locations := strings.Split (data, "_")
	if len (locations) > 32 {
		panic (invErr6)
	}

	for _, location := range locations {
		if location == "" {
			panic (invErr1)
		}
		locationData := strings.Split (location, "-")

		if len (locationData) == 1 {
			panic (invErr2)
		}

		if locationData [0] == "" {
			panic (invErr8)
		}

		if len (locationData) > 33 {
			panic (invErr7)
		}

		for index := 1; index <= len (locationData) - 1; index ++ {
			if dayMonthYear.Match (locationData [index]) == false {
				panic (invErr3)
			}
		}
	}
	// ...2... }

	// Validating existence of all locations. ...2... {
	// Constructing query required for validation. ...3... { ...
	query := `
		SELECT COUNT (id)
		FROM location
		WHERE id IN (?` + strings.Repeat (", ?", len (locations) - 1) + ")"
	// ...3... }

	errX := db.Ping ()
	if errX != nil {
		err_ := err.New (oprErr0.Error (), oprErr0.Class (), oprErr0.Type (), errX)
		panic (err_)
	}

	var noOfValidLocations int
	errY := db.QueryRow (query, locations [1:] ...).Scan (&noOfValidLocations)
	if errY != nil {
		err_ := err.New (oprErr1.Error (), oprErr1.Class (), oprErr1.Type (), errY)
		panic (errr_)
	}

	if noOfValidLocations != len (locations) {
		panic (oprErr2)
	}
	// ...2... }

	return validatedRequestData {data}
}

type validatedRequestData struct {
	data string
}

// fetchRecords () fetches all location state records matching the request of the user.
func (d *validatedRequestData) fetchRecords () (*requestRecords) {
	// Function definitions. ...1... {
	extractLocationIDs := func (requestData string) ([]string) {
		ids := []string {}
		locations := strings.Split (requestData, "_")
		for _, locationData := range locations {
			data := strings.Split (locationData, "-")
			ids = append (ids, data [0])
		}
		return ids
	}
	// ...1... }

	// Constructing query required to retrieve the sensor IDs of all locations. ...1... {
	queryB := `
		SELECT id, sensor
		FROM location
		WHERE id IN (?
	` + strings.Repeat (", ?", len (locations) - 1) + ")"
	// ...1... }

	// Retrieving the sensor IDs of all locations. ...1... {
	resultSet, errZ := db.Query (query, extractLocationIDs (d.data)...)
	if errZ != nil {
		err_ := err.New (oprErr3.Error (), oprErr3.Class (), oprErr3.Type (), errZ)
		panic (err_)
	}

	sensors := []string {}

	for resultSet.Next () {
		var (
			locationID string
			sensorID string
		)

		errA := resultSet.Scan (&locationID, &sensorID)
		if errA != nil {
			err_ := err.New (oprErr4.Error (), oprErr4.Class (), oprErr4.Type (), errA)
			panic (err_)
		}

		sensors = append (sensors, sensorID)
	}
	// ...1... }

	// Constructing query required to retrieve states from the database. ...1... {
	queryC := `
		SELECT UNIQUE state, day, time, sensor
		FROM state
		WHERE sensor IN (?
	` + strings.Repeat (", ?", len (sensors) - 1) + ")"
	// ...1... }

	// Retrieving the states of all locations. ...1... {
	var (
		states []_state
	)

	resultSetB, errB := db.Query (queryC, sensors...)
	if errB != nil {
		err_ := err.New (oprErr5.Error (), oprErr5.Class (), oprErr5.Type (), errB)
		panic (err_)
	}

	for resultSetB.Next () {
		someState := state {}

		errC := resultSetB.Scan (&someState.state, &someState.day, &someState.time, &someState.sensor)
		if errC != nil {
			err_ := err.New (oprErr6.Error (), oprErr6.Class (), oprErr6.Type (), errC)
			panic (err_)
		}

		states := append (states, someState)
	}
	// ...1... }

	return &requestRecords {states}	
}

type _state struct {
	state  string
	Day    string
	Time   string
	Sensor string
}

type requestRecords struct {
	states []_state
}

func (r *requestRecords) Organize () (result *organizedRequestRecords) {
	// Function definitions. ... {
	organizeByDay := func (sensorRecords []interface) (map[string] []_state) {
		records, errA := squaket.New (sensorRecords)
		if errA != nil {
			err_ := err.New (oprErr9.Error (), oprErr9.Class (), oprErr9.Type (), errA)
			panic (err_)
		}

		sensorsRecords, errB := records.Group ("Day")
		if errB != nil {
			err_ := err.New (oprErr10.Error (), oprErr10.Class (), oprErr10.Type (), errB)
			panic (err_)
		}

		organizedRecords := map[string] []_state {}

		iter := reflect.ValueOf (sensorsRecords).MapRange ()
		for iter.Next () {
			dayRecords := iter.Value ().Interface ().([]interface)
			sensorID := iter.Key ().Interface ().(string)

			stateTypeDayRecords := []_state {}
			for _, record := range dayRecords {
				stateTypeDayRecords = append (stateTypeDayRecords, record.(_state))
			}

			organizedRecords [sensorID] = stateTypeDayRecords
		}	

		return organizedRecords
	}
	// ... }

	records, errX := squaket.New (r.records)
	if errX != nil {
		err_ := err.New (oprErr7.Error (), oprErr7.Class (), oprErr7.Type (), errX)
		panic (err_)
	}

	sensorsRecords, errY := records.Group ("Sensor")
	if errY != nil {
		err_ := err.New (oprErr8.Error (), oprErr8.Class (), oprErr8.Type (), errY)
		panic (err_)
	}

	organizedRecords := organizedRequestRecords {
		map[string] map[string] []_states {}
	}

	iter := reflect.ValueOf (sensorsRecords).MapRange ()
	for iter.Next () {
		sensorRecords := iter.Value ().Interface ().([]interface)
		organizedSensorRecords := organizeByDay (sensorRecords)
		sensorID := iter.Key ().Interface ().(string)
		organizedRecords [sensorID] = organizedSensorRecords
	}

	return organizedRecords
}

type organizedRequestRecords struct {
	records map[string] map[string] []_state
}
