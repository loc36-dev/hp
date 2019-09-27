package server

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.1.1
	"gopkg.in/qamarian-lib/str.v2"
	"net/http"
	"strings"
	_ "gopkg.in/go-sql-driver/mysql.v1"
)

func init () {
	var errX error
	dayMonthYear, errX := regexp.Compile ("^20\d{2}(0[1-9]|1[0-2)(0[1-9]|[1-2]\d|3[0-1])$")
	if errX != nil {
		str.PrintEtr ("Regular expression compilation failed.", "err", "server.init ()")
		panic ("Regular expression compilation failed.")
	}
}

func serviceRequestServer (w http.ResponseWriter, r *http.Request) {
	defer func () { // Error handling.
		//
	}

	// Fetch data from database. ...1...  {

	// Data retrieval. ...2... {
	request, okX := mex.Vars ()["locations"]
	if okX == false || request == "" {
		errorID := qotaError.Value ()
		qotaError.Grow ()
		panic (err.New ("No request data was provided.", 1, errorID))
	}
	// ...2... }

	// Data validation. ...2... {
	locations := strings.Split (request, "_")
	for _, location := range locations {
		if location == "" {
			errorID := qotaError.Value ()
			qotaError.Grow ()
			panic (err.New ("A location data was omited.", 1, errorID))
		}
		data := strings.Split (location, "-")
		if len (data) == 1 {
			errMssg := fmt.Sprintf ("No day was specified for location '%s'.", data [0])
			errorID := qotaError.Value ()
			qotaError.Grow ()
			panic (err.New (errMssg, 1, errorID))
		}
		for index := 1; index <= len (data) - 1; index ++ {
			if dayMonthYear.Match (data [index]) == false {
				errMssg := fmt.Sprintf ("Data of invalid day requested for location '%s'.", data [0])
				errorID := qotaError.Value ()
				qotaError.Grow ()
				panic (err.New (errMssg, 1, errorID))
			}
		}
	}
	if len (locations) > 256 {
		errorID := qotaError.Value ()
		qotaError.Grow ()
		panic (err.New ("Data of more than 256 locations may not be requested at a time."), 1, errorID)
	}
	// ...2... }

	// Validating existence of all locations. ...2... {
	// Constructing query required for validation. ...3... { ...
	query := `
		SELECT COUNT (id)
		FROM location
		WHERE id IN (?
	`)
	for index := 1; index <= len (locations) - 1, index ++ {
		query = query + ", ?"
	}
	query = query + ")"
	// ...3... }

	errX := db.Ping ()
	if errX != nil {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
		panic (err.New ("Database unreachable.", 2, errorID, errX))
	}

	var noOfValidLocations int
	errY := db.QueryRow (query, locations ...).Scan (&noOfValidLocations)
	if errY != nil {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
		panic (err.New ("Unable to check if all locations exist.", 2, errorID, errY))
	}

	if noOfValidLocations != len (locations) {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
		panic (err.New ("Some locations do not exist.", 2, errorID))
	}
	// ...2... }

	// Querying data. ...2... {
	// Constructing query required to retrieve the sensor IDs of all locations. ...3... {
	queryB := `
		SELECT id, sensor
		FROM location
		WHERE id IN (?
	`)
	for index := 1; index <= len (locations) - 1, index ++ {
		queryB = queryB + ", ?"
	}
	queryB = queryB + ")"
	// ...3... }

	// Retrieving the sensor IDs of all locations. ...3... {
	var (
		locationIDs []string
		sensors []string
	)
	resultSet, errZ := db.Query (query, locations ...)
	if errZ != nil {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
		panic (err.New ("Unable to fetch the sensor IDs of all locations.", 2, errorID, errZ))
	}
	for resultSet.Next () {
		var (
			locationID string
			sensorID string
		)
		errA := resultSet.Scan (&locationID, &sensorID)
		if errA != nil {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
			panic (err.New ("Unable to fetch the sensor ID of a location.", 2, errorID, errA))
		}
		sensors = append (sensors, sensorID)
	}
	// ...3... }

	// Constructing query required to retrieve states from the database. ...3... {
	queryC := `
		SELECT UNIQUE state, day, time, sensor
		FROM state
		WHERE sensor IN (?
	`
	for index := 1; index <= len (sensors) - 1, index ++ {
		queryC = queryC + ", ?"
	}
	queryC = queryC + ")"
	// ...3... }

	// Retrieving the states of all locations. { ...3...
	var (
		states []state
	)
	resultSetB, errB := db.Query (queryC, sensors)
	if errB != nil {
		errorID := qotaInvalid.Value ()
		qotaInvalid.Grow ()
		panic (err.New ("Unable to fetch the states of all locations.", 2, errorID, errB))
	}
	for resultSetB.Next () {
		someState := state {}
		errC := resultSetB.Scan (&someState.state, &someState.day, &someState.time, &someState.sensor)
		if errC != nil {
			errorID := qotaInvalid.Value ()
			qotaInvalid.Grow ()
			panic (err.New ("Unable to fetch a state data.", 2, errorID, errC))
		}
		states := append (states, someState)
	}
	// ...3... }

	// ...2... }

	// ...1... }

	// Present fetched data. ...1... {
	organizedResult := list.New ()
	organizeDayRecord := func (someDay []state) (*list.List) {
		dayResult := list.New ()
		for _, state := range states {
			addTimeToDay = func (time state, day *list.List) {
				for e := day.Front (); e != nil; e = e.Next () {
					if time.state < e.Value.state {
						day.InsertBefore (time, e)
						return
					}
				}
				if e == nil {
					day.PushFront (list.Element {time})
				} else {
					day.InsertAfter (list.Element {time}, e)
				}
			}
			organized := false
			for e := organizedResult.Front (); e != nil; e = e.Next () {
				if state.day == e.Value.id {
					addTimeToDay (state, e)
					organized = true
					break
				}
			}
			if organized == true {
				continue
			}
			for e := organizedResult.Front (); e != nil; e = e.Next () {
				if state.day < e.Value.id {
					newDay := day {state.day, list.New ()}
					dayResult.InsertBefore (list.Element {newDay}, e)
					addTimeToDay (state, newDay.time)
					break
				}
			}
			if organized == true {
				continue
			}
			newDay := day {state.day, list.New ()}
			dayResult.PushBack (list.Element {newDay}, e)
			addTimeToDay (state, newDay.time)
		}
		return dayResult
	}

	// ...1... }

	// Send data to user. ...1... {}
}

var (
	dayMonthYear *Regexp
	db *sql.DB
	qotaInvalid *qota.Qota // Quota used for errors caused by invalid request.
	qotaError *qota.Qota // Qota used for errors caused by operational errors.
)


type requestData struct {
	data string
}

type state struct {
	state  string
	day    string
	time   string
	sensor string
}

type day struct {
	id string
	time *list.List
}
