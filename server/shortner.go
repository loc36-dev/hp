package server

import (
	"encoding/json"
	"errors"
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.3.0
	"gopkg.in/qamarian-lib/str.v2"
	"gopkg.in/qamarian-lib/structs.v0" // v0.1.0
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

		states := append (states, _state_New (state, day, time, sensor))
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

func requestRecords_New (records []*_state) (output *requestRecords) {
	output = &requestRecords {records}
	return output
}

type requestRecords struct {
	value []*_state
}

func (r *requestRecords) group () (result *groupedRequestRecords) {
	// Function definitions. .. {
	organizeByDay := func (sensorRecords []interface {}) (map[string] []*_state, error) {
		sensorsRecords, errB := structs.Group ("Day", sensorRecords...)
		if errB != nil {
			errMssg := "Unable to group records of the sensor."
			return nil, err.New (errMssg, 0, 0, errB)
		}

		organizedRecords := map[string] []*_state {}

		iter := reflect.ValueOf (sensorsRecords).MapRange ()
		for iter.Next () {
			dayRecords := iter.Value ().Interface ().([]interface {})
			sensorID := iter.Key ().Interface ().(string)

			stateTypeDayRecords := []*_state {}
			for _, record := range dayRecords {
				stateTypeDayRecords = append (stateTypeDayRecords, record.(*_state))
			}

			organizedRecords [sensorID] = stateTypeDayRecords
		}

		return organizedRecords
	}
	// .. }

	sensorsRecords, errY := structs.Group ("Sensor", r.value...)
	if errY != nil {
		err_ := err.New (oprErr8.Error (), oprErr8.Class (), oprErr8.Type (), errY)
		panic (err_)
	}

	groupedRecords := groupedRequestRecords_New ()

	iter := reflect.ValueOf (sensorsRecords).MapRange ()
	for iter.Next () {
		sensorRecords := iter.Value ().Interface ().([]interface {})
		organizedSensorRecords, errQ := organizeByDay (sensorRecords)
		if errQ != nil {
			err_ := err.New (oprErr10.Error (), oprErr10.Class (), oprErr10.Type (), errQ)
			panic (err_)
		}
		sensorID := iter.Key ().Interface ().(string)
		groupedRecords.add (sensorID, organizedSensorRecords)
	}

	return groupedRecords
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
	return s.day
}

func (s *_state) time () (string) {
	return s.time
}

func (s *_state) sensor () (string) {
	return s.sensor
}

// -- Boundary -- //

func groupedRequestRecords_New () (*groupedRequestRecords) {
	return &groupedRequestRecords {map[string] map[string][]*_state {}}
}

type groupedRequestRecords struct {
	value map[string] map[string] []_state
}

func (r *groupedRequestRecords) addSensorRecords (sensorID string, record map[string][]*_state) {
	r.value [sensorID] = record
}

func (r *groupedRequestRecords) complete () (*completeData) {
	// Function definitions. ..1.. {
	completeDays := func (days map[interface {}] []interface {}) (map[string][1440]*_pureState) {
		day := map[string][1440]_pureState {}

		iter := relect.ValueOf (day).MapRange ()
		for iter.Next () {
			dayID := iter.Key ().(string)
			dayStates := iter.Value ().([]interface {})

			pureStates := [1440]*_pureState {}
			for index, _ := range pureStates {
				pureStates [index] = _pureState_New (-1)
			}

			for _, value := range dayStates {
				if value.(*_state).time () == "0000" {
					continue
				}

				hour, _ := strconv.Atoi (value.(*_state).time ()[0:2])
				min,  _ := strconv.Atoi (value.(*_state).time ()[2:4])
				min = (hour * 60) + min

				minIndex := min - 1

				if (minIndex - 4) >= 0 && pureStates [minIndex - 4].state () == -1 {
					pureStates [minIndex - 4].set (byte (strconv.Atoi (value.(*_state).state ())))
				}
				if (minIndex - 3) >= 0 && pureStates [minIndex - 3].state () == -1 {
					pureStates [minIndex - 3].set (byte (strconv.Atoi (value.(*_state).state ())))
				}
				if (minIndex - 2) >= 0 && pureStates [minIndex - 2].state () == -1 {
					pureStates [minIndex - 2].set (byte (strconv.Atoi (value.(*_state).state ())))
				}
				if (minIndex - 1) >= 0 && pureStates [minIndex - 1].state () == -1 {
					pureStates [minIndex - 1].set (byte (strconv.Atoi (value.(*_state).state ())))
				}

				pureStates [minIndex].set (byte (strconv.Atoi (value.(_state).state ())))
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
	return &completeData {map[string] map[string][1440]*_pureState {}}
}

type completeData struct {
	value map[string] map[string][1440]*_pureState
}

func (d *completeData) add (sensorID string, data map[string][1440]*_pureState) {
	d.value [sensorID] = data
}

func (d *completeData) format () (*formattedData) {
	// Function definitions. ..1.. {
	formatDays := func (days map[string][1440]*_pureState) (map[string] []_formattedState) {
		formattedDays := map[string][]*_formattedState {}

		iter := reflect.ValueOf (days).MapRange ()
		for iter.Next () {
			formattedStates := []*_formattedState {}

			dayID := iter.Key.(string)
			dayStates := day [dayID]
			currentState, _ := dayStates[0].state ()

			for index, value := range dayStates {
				if currentState != value.state () {
					currentState = value.state ()

					hourMin :=  (index + 1) % 60
					hour := ((index + 1) - hourMin) / 60
					hr := str.PrependTillN (strconv.Itoa (hour), "0", 2)
					mn := str.PrependTillN (strconv.Itoa (hourMin), "0", 2)
					time := fmt.Sprintf ("%s%s", hr, min)

					someState := _formattedState_New (value.state (), time)
					formattedStates = append (formattedStates, someState)
				}
			}

			lastRecord := formattedStates [len (formattedStates) - 1]
			if lastRecord.endTime () != "2400" {
				someState := _formattedState_New (currentState, "2400")
				formattedStates = append (formattedStates, someState)
			}

			formattedDays [dayID] = formattedStates
		}

		return formattedDays
	}
	// ..1.. }

	data := formattedData_New ()

	iter := reflect.ValueOf (d.value).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().(string)
		sensorData := formatDays (d.value [sensorID])
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

func (s *_pureState) set (state byte) {
	*s = state
}

func (s *_pureState) state () (byte) {
	return byte (s)
}

// -- Boundary -- //

func formattedData_New () (*formattedData) {
	return &map[string] map[string][]*_formattedState {}
}

type formattedData map[string] map[string][]*_formattedState

func (d *formattedData) addSensor (sensorID string, record map[string][]*_formattedState) {
	*d [sensorID] = record
}

func (d *formattedData) marshal () (output *marshalledData) {
	output = marshalledData_New ()

	iter := reflect.ValueOf (d).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().(string)
		sensorData := iter.Key ().(map[string][]*_formattedState)
		data, errX := json.Marshal (sensorData)
		if errX != nil {
			err_ := err.New (oprErr11.Error (), oprErr11.Class (), oprErr11.Type (), errX)
			panic (err_)
		}
		output.addSensorData (sensorID, string (data))
	}

	return output
}

func _formattedState_New (state byte, endTime string) (*_formattedState) {
	return &_formattedState	{state, endTime}
}

type _formattedState struct {
	State byte
	EndTime string
}

func (s *_formattedState) state () (byte) {
	return r.State
}

func (s *_formattedState) endTime () (string) {
	return r.EndTime
}

// -- Boundary -- //

func marshalledData_New () (*marshalledData) {
	return &marshalledData {}
}

type marshalledData map[string]string

func (d *marshalledData) addSensorData (sensorID, data string) {
	*d [sensorID] = data
}

func (d *marshalledData) getSensorData (sensorID string) (string, bool) {
	return *d [sensorID]
}
