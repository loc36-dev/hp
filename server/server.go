package server

import (
	"fmt"
	"gopkg.in/qamarian-dtp/err.v0" // 0.4.0
	errLib "gopkg.in/qamarian-lib/err.v0" // 0.4.0
	"math/big"
	"net/http"
)

func serviceRequestServer (w http.ResponseWriter, r *http.Request) {
	// Error handling. ...1... {
	defer func () {
		panicReason := recover ()
		if panicReason == nil {
			return
		}

		reason, okX := panicReason ().(err.Error)

		if okX == true && reason.Class ().Cmp (invErrID) == 0 {
			errDetails := fmt.Sprintf ("Invalid request data. [%s]", errLib.Fup (reason))
			w.Write ([]byte (fmt.Sprintf (outputFormat, "rsp1", errDetails, "")))

			return
		} else {
			w.Write ([]byte (fmt.Sprintf ("An error occured.", "rsp2", errDetails, "")))

			errDetails := fmt.Sprintf ("A panic occured. [%v]", panicReason)
			if okX == true {
				errDetails = fmt.Sprintf ("An error occured. [%s]", errLib.Fup (reason))
			}
			log := log.New (stdErr, "", log.LDate | log.Ltime)
			log.WriteString (errDetails)

			return
		}
	} ()
	// ...1... }

	databaseRecords := new_requestData (r).fetchRecords (r)
	gData := databaseRecords.group ()
	cData := gData.complete ()
	fData := cData.format ()
	mData := fData.marshal ()

	// Present organized records. ...1... {
	//|--
	locationsIDs := extractLocationIDs (r)
	sensors, errX := locationsSensors (locationsIDs)
	if errX != nil {
		err_ := err.New (oprErr3.Error (), oprErr3.Class (), oprErr3.Type (), errX)
		panic (err_)
	}

	data := ""
	//--|

	for i, locationID := range locationsIDs {
		//|--
		comma := ","
		if i == len (locationsIDs) - 1 {
			comma = ""
		}
		//--|

		locationSensor, _ := sensors.getLocationSensor (locationID)
		locationData, _ := mData.getSensorData (locationSensor)

		locationData = fmt.Sprintf (`
			"%s": %s%s
			`, locationID, locationData, comma)

		data = data + locationData
	}

	output := fmt.Sprintf (outputFormat, "rsp0", "", data)
	fmt.Fprintf (w, output)
	// ...1... }
}

var (
	stdErr io.Writer = os.Stderr
	conf string
	outputFormat string = `{
Response: "%s",
Details: "%s"
Data: {
%s
}
}`
)

func init () {}
