package server

import (
	"gopkg.in/qamarian-dtp/qota.v0" // v0.
	"gopkg.in/qamarian-dtp/err.v0"
	"math/big"
)

var (
// Errors class ID
	invErrID *big.Int // Class ID for service errors caused by invalid request.
	oprErrID *big.Int // Class ID for service errors caused by operational errors.


// Errors iota
	iotaInv *qota.Qota // Iota used for service errors caused by invalid request.
	iotaOpr *qota.Qota // Iota used for service errors caused by operational errors.

// Invalid request data Errors

	invErr0 *err.Error
	invErr1 *err.Error
	invErr2 *err.Error
	invErr3 *err.Error
	invErr4 *err.Error
	invErr5 *err.Error
	invErr6 *err.Error
	invErr7 *err.Error
	invErr8 *err.Error
	invErr9 *err.Error

// Operational Errors

	oprErr0 *err.Error
	oprErr1 *err.Error
	oprErr2 *err.Error
	oprErr3 *err.Error
	oprErr4 *err.Error
	oprErr5 *err.Error
	oprErr6 *err.Error
	oprErr7 *err.Error
	oprErr8 *err.Error
	oprErr9 *err.Error
	oprErr10 *err.Error
	oprErr11 *err.Error

)

func init () {
	invErrID = big.NewInt (0)
	oprErrID = big.NewInt (1)

	iotaInv = qota.New ()
	iotaOpr = qota.New ()

	invErr0 = err.New ("No request data was provided.", invErrID, iotaInv.Value ())
	invErr1 = err.New ("A location data was omited.", invErrID, iotaInv.Value ())
	invErr2 = err.New ("No day was specified for a location.", invErrID, iotaInv.Value ())
	invErr3 = err.New ("Data of an invalid day requested for a location.", invErrID, iotaInv.Value ())
	invErr4 = err.New ("Data of more than 32 locations may not be requested at a time.", invErrID, iotaInv.Value ())
	invErr5 = err.New ("Request data too long.", invErrID, iotaInv.Value ())
	invErr6 = err.New ("Data of over 32 locations may not be requested at a tme.", invErrID, iotaInv.Value ())
	invErr7 = err.New ("Records of over 32 days may not be requested for a location, at a time.", invErrID, iotaInv.Value ())
	invErr8 = err.New ("The ID of a location was not provided.", invErrID, iotaInv.Value ())
	invErr9 = err.New ("Request data validation failed.", invErrID, iotaInv.Value ())

	oprErr0 = err.New ("Database unreachable.", oprErrID, iotaOpr.Value ())
	oprErr1 = err.New ("Unable to check if all locations exist.", oprErrID, iotaOpr.Value ())
	oprErr2 = err.New ("Some locations do not exist.", oprErrID, iotaOpr.Value ())
	oprErr3 = err.New ("Unable to fetch the sensor IDs of all locations.", oprErrID, iotaOpr.Value ())
	oprErr4 = err.New ("Unable to fetch the sensor ID of a location.", oprErrID, iotaOpr.Value ())
	oprErr5 = err.New ("Unable to fetch the states of all locations.", oprErrID, iotaOpr.Value ())
	oprErr6 = err.New ("Unable to fetch a state data.", oprErrID, iotaOpr.Value ())
	oprErr7 = err.New ("Unable to make request records a squaket.", oprErrID, iotaOpr.Value ())
	oprErr8 = err.New ("Unable to organize request records, based on sensor.", oprErrID, iotaOpr.Value ())
	oprErr9 = err.New ("Unable to make records of a sensor a squaket.", oprErrID, iotaOpr.Value ())
	oprErr10 = err.New ("Unable to organize records of a sensor, based on day.", oprErrID, iotaOpr.Value ())
	oprErr11 = err.New ("A sensor record could not be marshalled.", oprErrID, iotaOpr.Value ())
}
