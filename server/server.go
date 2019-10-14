package server

import (
	"fmt"
	"gopkg.in/qamarian-dtp/err.v0" // 0.4.0
	errLib "gopkg.in/qamarian-lib/err.v0" // 0.4.0
	"net/http"
	"../lib"
)

func ServiceRequestServer (w http.ResponseWriter, r *http.Request) {
	// Error handling. ..1.. {
	defer func () {
		panicReason := recover ()
		if panicReason == nil {
			return
		}

		reason, okX := panicReason.(err.Error)

		if okX == true && reason.Class ().Cmp (lib.InvErrID) == 0 {
			errDetails := fmt.Sprintf ("Invalid request data. [%s]", errLib.Fup (&reason))
			w.Write ([]byte (fmt.Sprintf (outputFormat, "rsp1", errDetails, "")))

			return
		} else {
			w.Write ([]byte (fmt.Sprintf (outputFormat, "rsp2", "An error occured.", "")))

			errDetails := fmt.Sprintf ("A panic occured. [%v]", panicReason)
			if okX == true {
				errDetails = fmt.Sprintf ("An error occured. [%s]", errLib.Fup (&reason))
			}
			lib.StdErr.Write ([]byte (errDetails))

			return
		}
	} ()
	// ..1.. }

	databaseRecords := requestData_New (r).fetchRecords ()
	gData := databaseRecords.group ()
	cData := gData.complete ()
	fData := cData.format ()
	mData := fData.marshal ()

	// Present organized records. ..1.. {
	//|--
	locationsIDs := extractLocationIDs (r)
	sensors, errX := locationsSensors (locationsIDs)
	if errX != nil {
		err_ := err.New (lib.OprErr3.Error (), lib.OprErr3.Class (), lib.OprErr3.Type (), errX)
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
	outputFormat string = `{

ServiceID: "0",
Version: "0.1.0",
Response: "%s",
Details: "%s",
Data: {
%s
}

}`
)
