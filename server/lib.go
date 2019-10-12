package server

import (
	"database/sql"
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	_ "gopkg.in/go-sql-driver/mysql.v1"
	"gopkg.in/qamarian-dtp/err.v0" // v0.4.0
	errLib "gopkg.in/qamarian-lib/err.v0" // v0.4.0
	"gopkg.in/qamarian-lib/str.v2" // v3.0.0
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"net/url"
	"os"
	"../lib"
)

// -- Boundary -- //

var (
	dayMonthYear *regexp.Regexp // Cache
	db *sql.DB           // Cache
)

func init () {
	// dayMonthYear initialization. ..1.. {
	var errX error
	dayMonthYear, errX = regexp.Compile (`^20\d{2}(0[1-9]|1[0-2)(0[1-9]|[1-2]\d|3[0-1])$`)
	if errX != nil {
		errY := err.New ("Regular expression compilation failed.", nil, nil, errX)
		str.PrintEtr (errLib.Fup (errY), "err", "init ()")
		os.Exit (1)
	}
	// ..1..}

	// db initialization. ..1.. {
	connURLFormat := "%s:%s@tcp(%s:%s)/%s?tls=skip-verify&serverPubKey=%s&timeout=%ss&writeTimeout=%ss&readTimeout=%ss"

	conf, errA := lib.NewConf ()
	if errA != nil {
		errB := err.New ("Unable to load service's configuration.", nil, nil, errA)
		str.PrintEtr (errLib.Fup (errB), "err", "init ()")
		os.Exit (1)
	}

	connURL := fmt.Sprintf (connURLFormat, url.QueryEscape ((*conf) ["username"]), url.QueryEscape ((*conf) ["pass"]), url.QueryEscape ((*conf) ["dbms_addr"]), url.QueryEscape ((*conf) ["dbms_port"]), url.QueryEscape ((*conf) ["db"]), url.QueryEscape ((*conf) ["conn_timeout"]), url.QueryEscape ((*conf) ["dbms_pub_key"]), url.QueryEscape ((*conf) ["write_timeout"]), url.QueryEscape ((*conf) ["read_timeout"]))

	var errC error
	db, errC = sql.Open ("mysql", connURL)
	if errC != nil {
		errD := err.New ("Database unreachable.", nil, nil, errC)
		str.PrintEtr (errLib.Fup (errD), "err", "init ()")
		os.Exit (1)
	}

	errE := db.Ping ()
	if errE != nil {
		errF := err.New ("Database unreachable.", nil, nil, errE)
		str.PrintEtr (errLib.Fup (errF), "err", "init ()")
		os.Exit (1)
	}
	// ..1.. }
}

// -- Boundary -- //

func extractLocationIDs (r *http.Request) ([]string) {
	output := []string {}

	data, _ := mux.Vars (r)["locations"]
	locationsData := strings.Split (data, "_")

	for _, locationData := range locationsData {
		segments := strings.Split (locationData, "-")
		output = append (output, segments [0])
	}

	return output
}

// -- Boundary -- //

func locationsSensors (locations []string) (*_locationsSensors, error) {
	query := `SELECT id, sensor
	FROM location
	WHERE id IN (?` + strings.Repeat (", ?", len (locations) - 1) + ")"

	errX := db.Ping ()
	if errX != nil {
		errMssg := "Database unreachable."
		return nil, err.New (errMssg, nil, nil, errX)
	}

	var (
		location string
		sensor string
	)

	// | --
	interfacedData := []interface {}{}
	for _, value := range locations {
		interfacedData := append (interfacedData, value)
	}
	// -- |
	result, errY := db.Query (query, interfacedData...)
	if errY != nil {
		errMssg := "Unable to successfully query database for locations sensors."
		return nil, err.New (errMssg, nil, nil, errY)
	}

	output := _locationsSensors_New ()

	for result.Next () {
		errZ := result.Scan (&location, &sensor)
		if errZ != nil {
			errMssg := "Unable to fetch the sensor of a location."
			return nil, err.New (errMssg, nil, nil, errZ)
		}

		output.add (location, sensor)
	}

	return output, nil
}

func _locationsSensors_New () (*_locationsSensors) {
	return &_locationsSensors {}
}

type _locationsSensors map[string]string

func (l *_locationsSensors) add (location, sensor string) {
	(*l) [location] = sensor
}

func (l *_locationsSensors) getLocationSensor (location string) (string, bool) {
	output, okX := (*l) [location]
	return output, okX
}

func (l *_locationsSensors) sensors () ([]string) {
	iter := reflect.ValueOf (l).MapRange ()
	output := []string {}

	for iter.Next () {
		sensorID := iter.Value ().Interface ().(string)
		output = append (output, sensorID)
	}

	return output
}
// -- Boundary -- //
