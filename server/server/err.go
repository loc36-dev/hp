package server

var (
	qotaInv *qota.Qota   // Qota used for service errors caused by invalid request.
	qotaOpr *qota.Qota   // Qota used for service errors caused by operational errors.
)

// Invalid request data Error 0 - 4

invErr0 := err.New ("No request data was provided.", 0, qotaInv.Value ())
qotaInv.Grow ()

invErr1 := err.New ("A location data was omited.", 0, qotaInv.Value ())
qotaInv.Grow ()

invErr2 := err.New ("No day was specified for a location.", 0, qotaInv.Value ())
qotaInv.Grow ()

invErr3 := err.New ("Data of invalid day requested for a location.", 0, qotaInv.Value ())
qotaInv.Grow ()

invErr4 := err.New ("Data of more than 146 locations may not be requested at a time."), 0, qotaInv.Value ())
qotaInv.Grow ()

// Operational Error 0 - 4

oprErr0 := err.New ("Database unreachable.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr1 := err.New ("Unable to check if all locations exist.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr2 := err.New ("Some locations do not exist.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr3 := err.New ("Unable to fetch the sensor IDs of all locations.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr4 := err.New ("Unable to fetch the sensor ID of a location.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

// Operational Error 5 - 9

oprErr5 := err.New ("Unable to fetch the states of all locations.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr6 := err.New ("Unable to fetch a state data.", 1, qotaOpr.Value ())
qotaOpr.Grow ()
