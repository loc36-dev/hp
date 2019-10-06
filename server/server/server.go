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

			errDetails := fmt.Sprintf ("An error occured. [%s]", errLib.Fup (reason))
			if okX == false {
				errDetails = fmt.Sprintf ("A panic occured. [%v]", panicReason)
			}
			log := log.New (stdErr, "", log.LDate | log.Ltime)
			log.WriteString (errDetails)
		}
	} ()
	// ...1... }

	databaseRecords := new_requestData (r).fetchRecords (r)
	oData := databaseRecords.organize ()
	cData := oData.complete ()
	fData := cData.format ()
	mData := fData.marshal ()

	// Present organized records. ...1... {
	location := extractLocationIDs (r)
	sensors, errX := locationsSensors (locations)
	if errX != nil {
		err_ := err.New (oprErr3.Error (), oprErr3.Class (), oprErr3.Type (), errX)
		panic (err_)
	}

	data := ""

	for i, location := range locations {
		comma := ","
		if i == len (locations) - 1 {
			comma = ""
		}

		locationSensor, _ := sensors.getLocationSensor (location)
		locationData, _ := mData.getSensorData (locationSensor)

		marshalled := fmt.Sprintf (`
		"%s": %s%s
		`, location, locationData, comma)

		data = data + locationData
	}

	output = fmt.Sprintf (output, data)

	fmt.Fprintf (w, output)
	// ...1... }
}

var (
	stdErr io.Writer = os.Stderr
	outputFormat string = `{
Response: "rsp0",
Details: "%s"
Data: {

%s

}
}`
)
