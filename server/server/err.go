package server

import (
	"gopkg.in/qamarian-dtp/qota.v0" // v0.
	"gopkg.in/qamarian-dtp/err.v0"
	"math/big"
)

func

var (
// Error IDs
	invErrID *big.Int = big.NewInt (0) //
	oprErrID *big.Int = big.NewInt (0) //

// Error iotas
	qotaInv *qota.Qota  = qota.New () // Qota used for service errors caused by invalid request.
	qotaOpr *qota.Qota  = qota.New () // Qota used for service errors caused by operational errors.

// Invalid request data Error 0 - 4

	invErr0 err.Error = err.New ("No request data was provided.", 0, uint8 (qotaInv.Value ()))
	invErr1 err.Error = err.New ("A location data was omited.", 0, qotaInv.Value ())
	invErr2 err.Error = err.New ("No day was specified for a location.", 0, qotaInv.Value ())
	invErr3 err.Error = err.New ("Data of an invalid day requested for a location.", 0, qotaInv.Value ())
	invErr4 err.Error = err.New ("Data of more than 32 locations may not be requested at a time.", 0, qotaInv.Value ())

// Invalid request data Error 5 - 9

	invErr5 err.Error = err.New ("Request data too long.", 0, qotaInv.Value ())
	invErr6 err.Error = err.New ("Data of over 32 locations may not be requested at a tme.", 0, qotaInv.Value ())
	invErr7 err.Error = err.New ("Records of over 32 days may not be requested for a location, at a time.", 0, qotaInv.Value ())
	invErr8 err.Error = err.New ("The ID of a location was not provided.", 0, qotaInv.Value ())
	invErr9 err.Error = err.New ("Request data validation failed.", 0, qotaInv.Value ())

// Operational Error 0 - 4

	oprErr0 err.Error = err.New ("Database unreachable.", 1, qotaOpr.Value ())
	oprErr1 err.Error = err.New ("Unable to check if all locations exist.", 1, qotaOpr.Value ())
	oprErr2 err.Error = err.New ("Some locations do not exist.", 1, qotaOpr.Value ())
	oprErr3 err.Error = err.New ("Unable to fetch the sensor IDs of all locations.", 1, qotaOpr.Value ())
	oprErr4 err.Error = err.New ("Unable to fetch the sensor ID of a location.", 1, qotaOpr.Value ())

// Operational Error 5 - 9

	oprErr5 err.Error = err.New ("Unable to fetch the states of all locations.", 1, qotaOpr.Value ())
	oprErr6 err.Error = err.New ("Unable to fetch a state data.", 1, qotaOpr.Value ())
	oprErr7 err.Error = err.New ("Unable to make request records a squaket.", 1, qotaOpr.Value ())
	oprErr8 err.Error = err.New ("Unable to organize request records, based on sensor.", 1, qotaOpr.Value ())
	oprErr9 err.Error = err.New ("Unable to make records of a sensor a squaket.", 1, qotaOpr.Value ())

// Operational Error 10 - 14

	oprErr10 err.Error = err.New ("Unable to organize records of a sensor, based on day.", 1, qotaOpr.Value ())
	oprErr11 err.Error = err.New ("A sensor record could not be marshalled.", 1, qotaOpr.Value ())

)
