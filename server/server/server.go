package server
/*
import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"gopkg.in/qamarian-dtp/err.v0" // v0.1.1
	"gopkg.in/qamarian-lib/str.v2"
	"net/http"
	"strings"
	_ "gopkg.in/go-sql-driver/mysql.v1"
)
*/

/*
func init () {
	var errX error
	dayMonthYear, errX := regexp.Compile ("^20\d{2}(0[1-9]|1[0-2)(0[1-9]|[1-2]\d|3[0-1])$")
	if errX != nil {
		str.PrintEtr ("Regular expression compilation failed.", "err", "init ()")
		panic ("Regular expression compilation failed.")
	}
}
*/

func serviceRequestServer (w http.ResponseWriter, r *http.Request) {
	// Error handling. ...1... {
	defer func () {
		//
	}
	// ...1... }

	validatedData := requestData {}.Validate ()
	databaseRecords  := validatedData.FetchRecords ()
	organizedData := databaseRecords.Organize ()

	// Present organized records. ...1... {
	// ...1... }
}

/*
var (
	dayMonthYear *Regexp // Cache
	db *sql.DB           // Cache
)

type day struct {
	id string
	time *list.List
}
*/
