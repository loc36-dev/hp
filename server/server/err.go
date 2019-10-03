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

invErr3 := err.New ("Data of an invalid day requested for a location.", 0, qotaInv.Value ())
qotaInv.Grow ()

invErr4 := err.New ("Data of more than 32 locations may not be requested at a time."), 0, qotaInv.Value ())
qotaInv.Grow ()

// Invalid request data Error 5 - 9

invErr5 := err.New ("Request data too long."), 0, qotaInv.Value ())
qotaInv.Grow ()

invErr6 := err.New ("Data of over 32 locations may not be requested at a tme."), 0, qotaInv.Value ())
qotaInv.Grow ()

invErr7 := err.New ("Records of over 32 days may not be requested for a location, at a time."), 0, qotaInv.Value ())
qotaInv.Grow ()

invErr8 := err.New ("The ID of a location was not provided."), 0, qotaInv.Value ())
qotaInv.Grow ()

invErr9 := err.New ("Request data validation failed."), 0, qotaInv.Value ())
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

oprErr7 := err.New ("Unable to make request records a squaket.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr8 := err.New ("Unable to organize request records, based on sensor.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr9 := err.New ("Unable to make records of a sensor a squaket.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

// Operational Error 10 - 14

oprErr10 := err.New ("Unable to organize records of a sensor, based on day.", 1, qotaOpr.Value ())
qotaOpr.Grow ()

oprErr11 := err.New ("A sensor record could not be marshalled.", 1, qotaOpr.Value ())
qotaOpr.Grow ()
