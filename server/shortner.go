package server

import (
	"encoding/json"
	"errors"
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.3.0
	"gopkg.in/qamarian-dtp/squaket.v0" // v0.1.1
	"gopkg.in/qamarian-lib/str.v2"
	"net/http"
	"reflect"
	"strings"
)

func requestData_New (request string) (*requestData) {
	data, _ := mux.Var (request)["locations"]
	return &requestData {data}
}

type requestData struct {
	value string
}

// fetchRecords () fetches all location state records matching the request of the user.
func (d *requestData) fetchRecords (r *http.Request) (*requestRecords) {
	// Request data validation and retrieval. ..1.. {
	var errX error

	errX = _requestData_New (r).validate ()
	if errX != nil {
		err_ := err.New (oprErr9.Error (), oprErr9.Class (), oprErr9.Type (), errX)
		panic (err_)
	}
	// ..1.. }

	// Fetching sensor IDs of all locations. ..1.. {
	sensors, errY := locationsSensors (extractLocationIDs (r))
	if errY != nil {
		err_ := err.New (oprErr3.Error (), oprErr3.Class (), oprErr3.Type (), errY)
		panic (err_)
	}
	onlySensorIDs = sensors.sensors ()
	// ..1.. }

	// Fetching state records of sensors. ..1.. {
	// | --
	queryC := `
		SELECT UNIQUE state, day, time, sensor
		FROM state
		WHERE sensor IN (?
	` + strings.Repeat (", ?", len (onlySensorIDs) - 1) + ")"
	// -- |

	resultSetB, errB := db.Query (queryC, onlySensorIDs...)
	if errB != nil {
		err_ := err.New (oprErr5.Error (), oprErr5.Class (), oprErr5.Type (), errB)
		panic (err_)
	}

	var (
		states []*_state
		state string
		day string
		time string
		sensor string
	)

	for resultSetB.Next () {
		errC := resultSetB.Scan (&state, &day, &time, &sensor)
		if errC != nil {
			err_ := err.New (oprErr6.Error (), oprErr6.Class (), oprErr6.Type (), errC)
			panic (err_)
		}

		states := append (states, new__state (state, day, time, sensor))
	}
	// ..1.. }

	return requestRecords_New (states)
}

func _requestData_New (data string) (*_requestData) {
	return &_requestData {data}
}

type _requestData struct {
	value string
}

// validate () checks if the request data of the client is valid. If the request data is not valid, the client's request would not be served.
func (d *_requestData) validate (r *http.Request) (error) {
	// Checking if request data was properly formatted. ..1.. {
	if d.value == "" {
		return errors,New ("No request data was provided.")
	} else if len (d.value) > 1024 {
		return errors.New ("Request data too long.")
	}

	locations := strings.Split (d.value, "_")
	if len (locations) > 32 {
		return error.New ("Data of over 32 locations requested.")
	}

	for _, location := range locations {
		if location == "" {
			return errors.New ("ID and days for a location were not provided.")
		}
		locationData := strings.Split (location, "-")

		if len (locationData) == 1 {
			return errors.New ("Days for a location were not provided.")
		}

		if locationData [0] == "" {
			return errors.New ("ID of a location was not provided.")
		}

		if len (locationData) > 33 {
			return errors.New ("Data of over 32 days requested for a location.")
		}

		for index := 1; index <= len (locationData) - 1; index ++ {
			if ! dayMonthYear.Match (locationData [index]) {
				return errors.New ("Data of an invalid day requested for a location.")
			}
		}
	}
	// ..1.. }

	// Validating existence of all locations. ..1.. {
	sensors, errX := locationsSensors (extractLocationIDs (r))
	if errX != nil {
		return err.New ("Unable to confirm existence of all locations.", nil, nil, errX)
	}

	ids := extractLocationIDs (r)

	for _, id := range ids {
		_, okX := sensors.getLocationSensor (id)
		if okX == false {
			return err.New ("One or more locations do not exist.")
		}
	}
	// ..1.. }

	return nil
}

// -- Boundary -- //

func requestRecords_New (records []_state) (*requestRecords) {
	var output requestRecords = records
	return &output
}

type requestRecords []_state

func (r *requestRecords) organize () (result *organizedRequestRecords) {
	// Function definitions. .. {
	organizeByDay := func (sensorRecords []interface {}) (map[string] []_state, error) {
		records, errA := squaket.New (sensorRecords)
		if errA != nil {
			errMssg := "Unable to make sensor records a squaket."
			return nil, err.New (errMssg, 0, 0, errA)
		}

		sensorsRecords, errB := records.Group ("Day")
		if errB != nil {
			errMssg := "Unable to group records of the sensor."
			return nil, err.New (errMssg, 0, 0, errB)
		}

		organizedRecords := map[string] []_state {}

		iter := reflect.ValueOf (sensorsRecords).MapRange ()
		for iter.Next () {
			dayRecords := iter.Value ().Interface ().([]interface {})
			sensorID := iter.Key ().Interface ().(string)

			stateTypeDayRecords := []_state {}
			for _, record := range dayRecords {
				stateTypeDayRecords = append (stateTypeDayRecords, record.(_state))
			}

			organizedRecords [sensorID] = stateTypeDayRecords
		}

		return organizedRecords
	}
	// .. }

	records, errX := squaket.New (r)
	if errX != nil {
		err_ := err.New (oprErr7.Error (), oprErr7.Class (), oprErr7.Type (), errX)
		panic (err_)
	}

	sensorsRecords, errY := records.Group ("Sensor")
	if errY != nil {
		err_ := err.New (oprErr8.Error (), oprErr8.Class (), oprErr8.Type (), errY)
		panic (err_)
	}

	organizedRecords := organizedRequestRecords_New ()

	iter := reflect.ValueOf (sensorsRecords).MapRange ()
	for iter.Next () {
		sensorRecords := iter.Value ().Interface ().([]interface {})
		organizedSensorRecords, errQ := organizeByDay (sensorRecords)
		if errQ != nil {
			err_ := err.New (oprErr10.Error (), oprErr10.Class (), oprErr10.Type (), errQ)
			panic (err_)
		}
		sensorID := iter.Key ().Interface ().(string)
		organizedRecords.add (sensorID, organizedSensorRecords)
	}

	return organizedRecords
}

func _state_New (state, day, time, sensor string) (*_state) {
	return &_state {state, day, time, sensor}
}

type _state struct {
	State  string
	Day    string
	Time   string
	Sensor string
}

func (s *_state) state () (string) {
	return s.state
}

func (s *_state) day () (string) {
	return s.state
}

func (s *_state) time () (string) {
	return s.state
}

func (s *_state) sensor () (string) {
	return s.state
}

func _requestData_New (r *http.Request) (*_requestData, error) {
	var requestData _requestData
	requestData, _ := mux.Vars (r)["locations"]
	return &requestData, nil
}

// -- Boundary -- //

func organizedRequestRecords_New () (*organizedRequestRecords) {
	return &map[string] map[string] []_state
}

type organizedRequestRecords map[string] map[string] []_state

func (r *organizedRequestRecords) addSensorRecords (sensorID string, record map[string] []_state) {
	r [sensorID] = record
}

func (r *organizedRequestRecords) complete () (*completeData) {
	// Function definitions. ..1.. {
	completeDays := func (days map[interface {}] []interface {}) (map[string][1440]_pureState) {
		day := map[string][1440]_pureState {}

		iter := relect.ValueOf (day).MapRange ()
		for iter.Next () {
			dayID := iter.Key ().(string)
			dayStates := iter.Value ().([]interface {})

			pureStates := [1440]_pureState {}
			for index, _ := range pureStates {
				pureStates [index] = -1
			}

			for _, value := range dayStates {
				if value.(_state).time () == "0000" {
					continue
				}

				hour, _ := strconv.Atoi (value.(_state).time [0:2])
				min, _ := strconv.Atoi (value.(_state).time [2:4])
				sec := (hour * 60) + min

				secIndex := sec - 1

				if (secIndex - 4) > 0 && pureStates [secIndex - 4] == -1 {
					pureStates [secIndex - 4] = byte (strconv.Atoi (value.(_state).state ()))
				}
				if (secIndex - 3) > 0 && pureStates [secIndex - 3] == -1 {
					pureStates [secIndex - 3] = byte (strconv.Atoi (value.(_state).state ()))
				}
				if (secIndex - 2) > 0 && pureStates [secIndex - 2] == -1 {
					pureStates [secIndex - 2] = byte (strconv.Atoi (value.(_state),state ()))
				}
				if (secIndex - 1) > 0 && pureStates [secIndex - 1] == -1 {
					pureStates [secIndex - 1] = byte (strconv.Atoi (value.(_state).state ()))
				}
				pureStates [secIndex] = byte (strconv.Atoi (value.(_state).state ()))
			}
			day [dayID] = pureStates
		}

		return day
	}
	// .. }

	data := completeData_New ()
		
	iter := reflect.ValueOf (r).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().(string)
		sensorData := completeDays (iter.Value ().(map[interface {}] []interface {}))
		data.add (sensorID, sensorData)
	}

	return data
}

// -- Boundary -- //

func completeData_New () (*completeData) {
	return &map[string] map[string] [1440]*_pureState {}
}

type completeData map[string] map[string] [1440]*_pureState

func (d *completeData) add (sensorID string, data map[string] [1440]*_pureState) {
	d [sensorID] = data
}

func (d *completeData) format () (*formatedData) {
	// Function definitions. ..1.. {
	formatDays := func (days map[interface {}] []interface {}) (map[string] []_formattedState) {
		formattedDays := map[string] []_formattedState {}

		iter := reflect.ValueOf (days).MapRange ()
		for iter.Next () {
			dayID := iter.Key.(string)
			dayStates := iter.Value.([]interface {})

			formattedStates := []_formattedState {}

			currentState, _ := dayStates [0].(_pureState)

			for index, value := range dayStates {
				if currentState != value.(_pureState) {
					currentState = value.(_pureState)
					min :=  (index + 1) % 60
					hour := ((index + 1) - min) / 60
					time := fmt.Sprintf ("%s%s", str.PrependTillN (strconv.Itoa (hour), "0", 2),
						str.PrependTillN (strconv.Itoa (min), "0", 2))
					someState := _formattedState_New (value.(_pureState).state (), time)
					formattedStates = append (formattedStates, someState)
				}
			}

			if formattedStates [len (formattedStates) - 1].endTime () != "2400" {
				someState := _formattedState_New (formattedStates [len (formattedStates) - 1].state (), "2400")
				formattedStates = append (formattedStates, someState)
			}

			formattedDays [dayID] = formattedStates
		}

		return formattedDays
	}
	// ..1.. }

	data := &formatedData {
		map[string] map[string] []_formattedState {},
	}

	iter := reflect.ValueOf (r.records).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().(string)
		sensorData := formatDays (iter.Key ().(map[interface {}] []interface {}))
		data.addSensor (sensorID, sensorData)
	}

	return data
}

func _pureState_New (state byte) (*_pureState) {
	var pureState _pureState
	pureState = state
	return &pureState
}

type _pureState byte

func (s *_pureState) state () (byte) {
	return byte (s)
}

// -- Boundary -- //

func formattedData_New () (*formattedData) {
	return &map[string] map[string] []_formattedState
}

type formattedData map[string] map[string] []_formattedState

func (d *formattedData) addSensor (sensorID string, record map[string] []_formattedState) {
	d [sensorID] = record
}

func (d *formattedData) marshal () (output *marshalledData) {
	iter := reflect.ValueOf (d).MapRange ()
	output = marshalledData_New ()

	for iter.Next () {
		sensorID := iter.Key ().(string)
		sensorData := iter.Key ().(map[interface {}] []interface {})
		data, errX := json.Marshal (sensorData)
		if errX != nil {
			err_ := err.New (oprErr11.Error (), oprErr11.Class (), oprErr11.Type (), errX)
			panic (err_)
		}
		output.addSensorData (sensorID, string (data))
	}

	return output
}

func _formattedState_New (state int, endTime string) (*_formattedState) {}

type _formattedState struct {
	State int
	EndTime string
}

func (s *_formattedState) state () (int) {
	return r.State
}

func (s *_formattedState) endTime () (string) {
	return r.EndTime
}

// -- Boundary -- //

func marshalledData_New () (*marshalledData) {
	var output marshalledData
	return &output
}

type marshalledData map[string]string

func (d *marshalledData) addSensorData (sensorID, data string) {
	d [sensorID] = data
}

func (d *marshalledData) getSensorData (sensorID string) (string, bool) {
	return d [sensorID]
}
