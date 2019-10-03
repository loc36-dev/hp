package server

import (
	"fmt"
	"net/http"
)

func serviceRequestServer (w http.ResponseWriter, r *http.Request) {
	// Error handling. ...1... {
	defer func () {
		//
	}
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

	output := `{
Response: "rsp0",
Data: {

%s

}
}`

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
