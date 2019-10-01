package server

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
	requestResponse := organizedData.Process (validatedData)

	// Present organized records. ...1... {
	// ...1... }
}

var (
	dayMonthYear *Regexp // Cache
	db *sql.DB           // Cache
)
