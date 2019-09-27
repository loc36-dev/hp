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
	validatedData := requestData {}.Validate ()
	databaseRecords  := validatedData.FetchRecords ()
	organizedData := databaseRecords.Organize ()
	// Present organized data to user. ...1... {}
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
