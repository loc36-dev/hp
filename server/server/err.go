package server

var (
	qotaInv *qota.Qota   // Qota used for service errors caused by invalid request.
	qotaOpr *qota.Qota   // Qota used for service errors caused by operational errors.
)

// Invalid request data Error 1 - 5

invErr1 := err.New ("No request data was provided.", 1, qotaInv.Value ())
qotaInv.Grow ()

invErr2 := err.New ("A location data was omited.", 1, qotaInv.Value ())
qotaInv.Grow ()

invErr3 := err.New ("No day was specified for location a location.", 1, qotaInv.Value ())
qotaInv.Grow ()

invErr4 := err.New ("Data of invalid day requested for a location.", 1, qotaInv.Value ())
qotaInv.Grow ()

invErr5 := err.New ("Data of more than 256 locations may not be requested at a time."), 1, qotaInv.Value ())
qotaInv.Grow ()

// Operational Error 1 - 5

oprErr1 := err.New ("Database unreachable.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr2 := err.New ("Unable to check if all locations exist.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr3 := err.New ("Some locations do not exist.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr4 := err.New ("Unable to fetch the sensor IDs of all locations.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr5 := err.New ("Unable to fetch the sensor ID of a location.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

// Operational Error 9 - 10

oprErr6 := err.New ("Unable to fetch the states of all locations.", 2, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr7 := err.New ("Unable to fetch a state data.", 2, qotaOpr.Value ())
qotaOpr.Grow ()
