package lib

import (
	"gopkg.in/qamarian-dtp/qota.v0" // v0.2.0
	"gopkg.in/qamarian-dtp/err.v0" // 0.4.0
	"math/big"
)

var (
// Errors class ID
	InvErrID *big.Int // Class ID for service errors caused by invalid request.
	OprErrID *big.Int // Class ID for service errors caused by operational errors.


// Errors iota
	iotaInv *qota.Qota // Iota used for service errors caused by invalid request.
	iotaOpr *qota.Qota // Iota used for service errors caused by operational errors.


// Invalid request data Errors
	InvErr0 *err.Error
	InvErr1 *err.Error
	InvErr2 *err.Error
	InvErr3 *err.Error
	InvErr4 *err.Error
	InvErr5 *err.Error
	InvErr6 *err.Error
	InvErr7 *err.Error
	InvErr8 *err.Error
	InvErr9 *err.Error


// Operational Errors
	OprErr0 *err.Error
	OprErr1 *err.Error
	OprErr2 *err.Error
	OprErr3 *err.Error
	OprErr4 *err.Error
	OprErr5 *err.Error
	OprErr6 *err.Error
	OprErr7 *err.Error
	OprErr8 *err.Error
	OprErr9 *err.Error
	OprErr10 *err.Error
	OprErr11 *err.Error

)

func init () {
	InvErrID = big.NewInt (0)
	OprErrID = big.NewInt (1)

	iotaInv = qota.New ()
	iotaOpr = qota.New ()

	InvErr0 = err.New ("No request data was provided.", InvErrID, iotaInv.Value ())
	InvErr1 = err.New ("A location data was omited.", InvErrID, iotaInv.Value ())
	InvErr2 = err.New ("No day was specified for a location.", InvErrID, iotaInv.Value ())
	InvErr3 = err.New ("Data of an invalid day requested for a location.", InvErrID, iotaInv.Value ())
	InvErr4 = err.New ("Data of more than 32 locations may not be requested at a time.", InvErrID, iotaInv.Value ())
	InvErr5 = err.New ("Request data too long.", InvErrID, iotaInv.Value ())
	InvErr6 = err.New ("Data of over 32 locations may not be requested at a tme.", InvErrID, iotaInv.Value ())
	InvErr7 = err.New ("Records of over 32 days may not be requested for a location, at a time.", InvErrID, iotaInv.Value ())
	InvErr8 = err.New ("The ID of a location was not provided.", InvErrID, iotaInv.Value ())
	InvErr9 = err.New ("Request data validation failed.", InvErrID, iotaInv.Value ())

	OprErr0 = err.New ("Database unreachable.", OprErrID, iotaOpr.Value ())
	OprErr1 = err.New ("Unable to check if all locations exist.", OprErrID, iotaOpr.Value ())
	OprErr2 = err.New ("Some locations do not exist.", OprErrID, iotaOpr.Value ())
	OprErr3 = err.New ("Unable to fetch the sensor IDs of all locations.", OprErrID, iotaOpr.Value ())
	OprErr4 = err.New ("Unable to fetch the sensor ID of a location.", OprErrID, iotaOpr.Value ())
	OprErr5 = err.New ("Unable to fetch the states of all locations.", OprErrID, iotaOpr.Value ())
	OprErr6 = err.New ("Unable to fetch a state data.", OprErrID, iotaOpr.Value ())
	OprErr7 = err.New ("Unable to make request records a squaket.", OprErrID, iotaOpr.Value ())
	OprErr8 = err.New ("Unable to organize request records, based on sensor.", OprErrID, iotaOpr.Value ())
	OprErr9 = err.New ("Unable to make records of a sensor a squaket.", OprErrID, iotaOpr.Value ())
	OprErr10 = err.New ("Unable to organize records of a sensor, based on day.", OprErrID, iotaOpr.Value ())
	OprErr11 = err.New ("A sensor record could not be marshalled.", OprErrID, iotaOpr.Value ())
}
	