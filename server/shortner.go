package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.3.0
	"gopkg.in/qamarian-lib/str.v2"
	"gopkg.in/qamarian-lib/structs.v0" // v0.1.0
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"../lib"
)

func requestData_New (request *http.Request) (*requestData) {
	return &requestData {request}
}

type requestData struct {
	value *http.Request
}

// fetchRecords () fetches all location state records matching the request of the user.
func (d *requestData) fetchRecords () (*requestRecords) {
	// Request data validation and retrieval. ..1.. {
	errX := _requestData_New (d.value).validate ()
	if errX != nil {
		err_ := err.New (lib.InvErr9.Error (), lib.InvErr9.Class (), lib.InvErr9.Type (), errX)
		panic (err_)
	}
	// ..1.. }

	// Fetching sensor IDs of all locations. ..1.. {
	sensors, errY := locationsSensors (extractLocationIDs (d.value))
	if errY != nil {
		err_ := err.New (lib.OprErr3.Error (), lib.OprErr3.Class (), lib.OprErr3.Type (), errY)
		panic (err_)
	}
	onlySensorIDs := sensors.sensors ()
	// ..1.. }

	// Fetching state records of sensors. ..1.. {
	// | --
	queryC := `
		SELECT UNIQUE state, day, time, sensor
		FROM state
		WHERE sensor IN (?
	` + strings.Repeat (", ?", len (onlySensorIDs) - 1) + ")"
	// -- |

	// | --
	interfacedDataX := []interface {}{}
	for _, value := range onlySensorIDs {
		interfacedDataX = append (interfacedDataX, value)
	}
	// -- |
	resultSetB, errB := db.Query (queryC, interfacedDataX...)
	if errB != nil {
		err_ := err.New (lib.OprErr5.Error (), lib.OprErr5.Class (), lib.OprErr5.Type (), errB)
		panic (err_)
	}

	var (
		states []*_state
		state int8
		day string
		time string
		sensor string
	)

	for resultSetB.Next () {
		errC := resultSetB.Scan (&state, &day, &time, &sensor)
		if errC != nil {
			err_ := err.New (lib.OprErr6.Error (), lib.OprErr6.Class (), lib.OprErr6.Type (), errC)
			panic (err_)
		}

		states = append (states, _state_New (state, day, time, sensor))
	}
	// ..1.. }

	return requestRecords_New (states)
}

func _requestData_New (data *http.Request) (*_requestData) {
	return &_requestData {data}
}

type _requestData struct {
	value *http.Request
}

// validate () checks if the request data of the client is valid. If the request data is not valid, the client's request would not be served.
func (d *_requestData) validate () (error) {
	value, _ := mux.Vars (d.value)["locations"]

	// Checking if request data was properly formatted. ..1.. {
	if value == "" {
		return errors.New ("No request data was provided.")
	} else if len (value) > 1024 {
		return errors.New ("Request data too long.")
	}

	locations := strings.Split (value, "_")
	if len (locations) > 32 {
		return errors.New ("Data of over 32 locations requested.")
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
			if ! dayMonthYear.Match ([]byte (locationData [index])) {
				return errors.New ("Data of an invalid day requested for a location.")
			}
		}
	}
	// ..1.. }

	// Validating existence of all locations. ..1.. {
	sensors, errX := locationsSensors (extractLocationIDs (d.value))
	if errX != nil {
		return err.New ("Unable to confirm existence of all locations.", nil, nil, errX)
	}

	ids := extractLocationIDs (d.value)

	for _, id := range ids {
		_, okX := sensors.getLocationSensor (id)
		if okX == false {
			return err.New ("One or more locations do not exist.", nil, nil)
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
			return nil, err.New (errMssg, nil, nil, errB)
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

		return organizedRecords, nil
	}
	// .. }

	// | --
	interfacedDataX := []interface {}{}
	for _, value := range r.value {
		interfacedDataX = append (interfacedDataX, value)
	}
	// -- |
	sensorsRecords, errY := structs.Group ("Sensor", interfacedDataX...)
	if errY != nil {
		err_ := err.New (lib.OprErr8.Error (), lib.OprErr8.Class (), lib.OprErr8.Type (), errY)
		panic (err_)
	}

	groupedRecords := groupedRequestRecords_New ()

	iter := reflect.ValueOf (sensorsRecords).MapRange ()
	for iter.Next () {
		sensorRecords := iter.Value ().Interface ().([]interface {})
		groupedSensorRecords, errQ := organizeByDay (sensorRecords)
		if errQ != nil {
			err_ := err.New (lib.OprErr10.Error (), lib.OprErr10.Class (), lib.OprErr10.Type (), errQ)
			panic (err_)
		}
		sensorID := iter.Key ().Interface ().(string)
		groupedRecords.addSensorRecords (sensorID, groupedSensorRecords)
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
	return s.State
}

func (s *_state) day () (string) {
	return s.Day
}

func (s *_state) time () (string) {
	return s.Time
}

func (s *_state) sensor () (string) {
	return s.Sensor
}

// -- Boundary -- //

func groupedRequestRecords_New () (*groupedRequestRecords) {
	return &groupedRequestRecords {map[string] map[string][]*_state {}}
}

type groupedRequestRecords struct {
	value map[string] map[string][]*_state
}

func (r *groupedRequestRecords) addSensorRecords (sensorID string, record map[string][]*_state) {
	r.value [sensorID] = record
}

func (r *groupedRequestRecords) complete () (*completeData) {
	// Function definitions. ..1.. {
	completeDays := func (days map[interface {}] []interface {}) (map[string][1440]*_pureState) {
		day := map[string][1440]*_pureState {}

		iter := reflect.ValueOf (days).MapRange ()
		for iter.Next () {
			dayID := iter.Key ().Interface ().(string)
			dayStates := iter.Value ().Interface ().([]interface {})

			pureStates := [1440]*_pureState {}
			for index, _ := range pureStates {
				pureStates [index] = _pureState_New (2)
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
					intValue, _ := strconv.Atoi (value.(*_state).state ())
					pureStates [minIndex - 4].set (int8 (intValue))
				}
				if (minIndex - 3) >= 0 && pureStates [minIndex - 3].state () == -1 {
					intValue, _ := strconv.Atoi (value.(*_state).state ())
					pureStates [minIndex - 3].set (int8 (intValue))
				}
				if (minIndex - 2) >= 0 && pureStates [minIndex - 2].state () == -1 {
					intValue, _ := strconv.Atoi (value.(*_state).state ())
					pureStates [minIndex - 2].set (int8 (intValue))
				}
				if (minIndex - 1) >= 0 && pureStates [minIndex - 1].state () == -1 {
					intValue, _ := strconv.Atoi (value.(*_state).state ())
					pureStates [minIndex - 1].set (int8 (intValue))
				}

				intValue, _ := strconv.Atoi (value.(*_state).state ())
				pureStates [minIndex].set (int8 (intValue))
			}
			day [dayID] = pureStates
		}

		return day
	}
	// .. }

	data := completeData_New ()
		
	iter := reflect.ValueOf (r).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().Interface ().(string)
		sensorData := completeDays (iter.Value ().Interface ().(map[interface {}] []interface {}))
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
	formatDays := func (days map[string][1440]*_pureState) (map[string] []*_formattedState) {
		formattedDays := map[string][]*_formattedState {}

		iter := reflect.ValueOf (days).MapRange ()
		for iter.Next () {
			formattedStates := []*_formattedState {}

			dayID := iter.Key ().Interface ().(string)
			dayStates := days [dayID]
			currentState := dayStates[0].state ()

			for index, value := range dayStates {
				if currentState != value.state () {
					currentState = value.state ()

					hourMin :=  (index + 1) % 60
					hour := ((index + 1) - hourMin) / 60
					hr := str.PrependTillN (strconv.Itoa (hour), "0", 2)
					mn := str.PrependTillN (strconv.Itoa (hourMin), "0", 2)
					time := fmt.Sprintf ("%s%s", hr, mn)

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
		sensorID := iter.Key ().Interface ().(string)
		sensorData := formatDays (d.value [sensorID])
		data.addSensor (sensorID, sensorData)
	}

	return data
}

func _pureState_New (state int8) (*_pureState) {
	var pureState _pureState
	pureState = _pureState {state}
	return &pureState
}

type _pureState struct {
	value int8
}

func (s *_pureState) set (state int8) {
	(*s).value = state
}

func (s *_pureState) state () (int8) {
	return (*s).value
}

// -- Boundary -- //

func formattedData_New () (*formattedData) {
	return & formattedData {map[string] map[string][]*_formattedState {}}
}

type formattedData struct {
	value map[string] map[string][]*_formattedState
}

func (d *formattedData) addSensor (sensorID string, record map[string][]*_formattedState) {
	(*d).value [sensorID] = record
}

func (d *formattedData) marshal () (output *marshalledData) {
	output = marshalledData_New ()

	iter := reflect.ValueOf ((*d).value).MapRange ()
	for iter.Next () {
		sensorID := iter.Key ().Interface ().(string)
		sensorData := iter.Value ().Interface ().(map[string][]*_formattedState)
		data, errX := json.Marshal (sensorData)
		if errX != nil {
			err_ := err.New (lib.OprErr11.Error (), lib.OprErr11.Class (), lib.OprErr11.Type (), errX)
			panic (err_)
		}
		output.addSensorData (sensorID, string (data))
	}

	return output
}

func _formattedState_New (state int8, endTime string) (*_formattedState) {
	return &_formattedState	{state, endTime}
}

type _formattedState struct {
	State int8
	EndTime string
}

func (s *_formattedState) state () (int8) {
	return (*s).State
}

func (s *_formattedState) endTime () (string) {
	return (*s).EndTime
}

// -- Boundary -- //

func marshalledData_New () (*marshalledData) {
	return &marshalledData {map[string]string {}}
}

type marshalledData struct {
	value map[string]string
}

func (d *marshalledData) addSensorData (sensorID, data string) {
	(*d).value [sensorID] = data
}

func (d *marshalledData) getSensorData (sensorID string) (string, bool) {
	value, okX := (*d).value [sensorID]
	return value, okX
}
