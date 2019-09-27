package server

// Invalid request data Error 1 - 5

errorID := qotaError.Value ()
qotaError.Grow ()
invErr1 := err.New ("No request data was provided.", 1, errorID)

errorID := qotaError.Value ()
qotaError.Grow ()
invErr2 := err.New ("A location data was omited.", 1, errorID)

errorID := qotaError.Value ()
qotaError.Grow ()
invErr3 := err.New ("No day was specified for location a location.", 1, errorID)

errorID := qotaError.Value ()
qotaError.Grow ()
invErr4 := err.New ("Data of invalid day requested for a location.", 1, errorID)

errorID := qotaError.Value ()
qotaError.Grow ()
invErr5 := err.New ("Data of more than 256 locations may not be requested at a time."), 1, errorID)

// Operational Error 1 - 5

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr1 := err.New ("Database unreachable.", 2, errorID, errX)

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr2 := err.New ("Unable to check if all locations exist.", 2, errorID, errY)

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr3 := err.New ("Some locations do not exist.", 2, errorID)

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr4 := err.New ("Unable to fetch the sensor IDs of all locations.", 2, errorID, errZ)

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr5 := err.New ("Unable to fetch the sensor ID of a location.", 2, errorID, errA)

// Operational Error 9 - 10

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr6 := err.New ("Unable to fetch the states of all locations.", 2, errorID, errB)

errorID := qotaInvalid.Value ()
qotaInvalid.Grow ()
oprErr7 := err.New ("Unable to fetch a state data.", 2, errorID, errC)
