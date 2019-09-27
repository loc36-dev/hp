package server

type requestData struct {}

func (d *requestData) Validate () (*validatedRequestData) {
		// Data retrieval. ...2... {
	request, okX := mex.Vars ()["locations"]
	if okX == false || request == "" {
		panic (invErr1)
	}
	// ...2... }

	// Data validation. ...2... {
	locations := strings.Split (request, "_")
	for _, location := range locations {
		if location == "" {
			panic (invErr2)
		}
		data := strings.Split (location, "-")
		if len (data) == 1 {
			panic (invErr3)
		}
		for index := 1; index <= len (data) - 1; index ++ {
			if dayMonthYear.Match (data [index]) == false {
				panic (invErr4)
			}
		}
	}
	if len (locations) > 256 {
		panic (invErr5)
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
		panic (oprErr1)
	}

	var noOfValidLocations int
	errY := db.QueryRow (query, locations ...).Scan (&noOfValidLocations)
	if errY != nil {
		panic (oprErr2)
	}

	if noOfValidLocations != len (locations) {
		panic (oprErr3)
	}
	// ...2... }
}

type validatedRequestData struct {}

func (d *validatedRequestData) FetchRecords () (*requestRecords) {
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
		panic (oprErr4)
	}
	for resultSet.Next () {
		var (
			locationID string
			sensorID string
		)
		errA := resultSet.Scan (&locationID, &sensorID)
		if errA != nil {
			panic (oprErr5)
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
		panic (oprErr6)
	}
	for resultSetB.Next () {
		someState := state {}
		errC := resultSetB.Scan (&someState.state, &someState.day, &someState.time, &someState.sensor)
		if errC != nil {
			panic (oprErr7)
		}
		states := append (states, someState)
	}
	// ...3... }

	// ...2... }

}

type requestRecords struct {}

func (d *requestRecords) Organize () (*organizedRequestRecords) {
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
}

type organizedRequestRecords struct {}
