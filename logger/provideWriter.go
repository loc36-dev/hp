package logger

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	_ "gopkg.in/mattn/go-sqlite3.v1"
)

// provideWriter () returns the log file of the RbS. It fetches the logfile's path from an
// SQLite 3 database.
func provideWriter () (io.Writer, error) {
	db, errX := sql.Open ("sqlite3", "file:res/conf.db3")
	if errX != nil {
		errMssg := fmt.Sprintf ("RbS's configuration file could not be opened. " +
			"[%s]", errX.Error ())
		return nil, errors.New (errMssg)
	}
	query := `SELECT filepath
		FROM filepath
		WHERE name == ?
	`
	logfilePath := ""
	errY := db.QueryRow (query, "logfile").Scan (&logfilePath)
	if errY != nil {
		errMssg := fmt.Sprintf ("RbS's logfile path could not be retrieved " +
			"from the configuration file. [%s]", errY.Error ())
		return nil, errors.New (errMssg)
	}
	logfile, errZ := os.Open (logfilePath, O_WRONLY | O_APPEND | O_SYNC, 0660)
	if errZ != nil {
		errMssg := fmt.Sprintf ("RbS's logfile could not be opened. [%s]",
			errZ. Error ())
		return nil, errors.New (errMssg)
	}
	return logfile, nil
}
