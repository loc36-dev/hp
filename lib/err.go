package lib

import (
	"gopkg.in/qamarian-dtp/qota.v0" // v0.2.0
	"gopkg.in/qamarian-dtp/err.v0" // 0.4.0
	"math/big"
)

var (
// Errors class ID
	InvErrID func () (*big.Int) // Class ID for service errors caused by invalid request.
	OprErrID func () (*big.Int) // Class ID for service errors caused by operational errors.


// Errors iota
	iotaInv *qota.Qota // Iota used for service errors caused by invalid request.
	iotaOpr *qota.Qota // Iota used for service errors caused by operational errors.


// Invalid request data Errors
	InvErr0 func () (*err.Error)
	InvErr1 func () (*err.Error)
	InvErr2 func () (*err.Error)
	InvErr3 func () (*err.Error)
	InvErr4 func () (*err.Error)
	InvErr5 func () (*err.Error)
	InvErr6 func () (*err.Error)
	InvErr7 func () (*err.Error)
	InvErr8 func () (*err.Error)
	InvErr9 func () (*err.Error)


// Operational Errors
	OprErr0 func () (*err.Error)
	OprErr1 func () (*err.Error)
	OprErr2 func () (*err.Error)
	OprErr3 func () (*err.Error)
	OprErr4 func () (*err.Error)
	OprErr5 func () (*err.Error)
	OprErr6 func () (*err.Error)
	OprErr7 func () (*err.Error)
	OprErr8 func () (*err.Error)
	OprErr9 func () (*err.Error)
	OprErr10 func () (*err.Error)
	OprErr11 func () (*err.Error)

)

func init () {
	InvErrID = func () (*big.Int) { return big.NewInt (0) }
	OprErrID = func () (*big.Int) { return big.NewInt (1) }

	iotaInv = qota.New ()
	iotaOpr = qota.New ()

	InvErr0 = func () (*err.Error) { return err.New ("No request data was provided.", InvErrID (), iotaInv.Value ()))
	InvErr1 = func () (*err.Error) { return err.New ("A location data was omited.", InvErrID (), iotaInv.Value ()))
	InvErr2 = func () (*err.Error) { return err.New ("No day was specified for a location.", InvErrID (), iotaInv.Value ()))
	InvErr3 = func () (*err.Error) { return err.New ("Data of an invalid day requested for a location.", InvErrID (), iotaInv.Value ()))
	InvErr4 = func () (*err.Error) { return err.New ("Data of more than 32 locations may not be requested at a time.", InvErrID (), iotaInv.Value ()))
	InvErr5 = func () (*err.Error) { return err.New ("Request data too long.", InvErrID (), iotaInv.Value ()))
	InvErr6 = func () (*err.Error) { return err.New ("Data of over 32 locations may not be requested at a tme.", InvErrID (), iotaInv.Value ()))
	InvErr7 = func () (*err.Error) { return err.New ("Records of over 32 days may not be requested for a location, at a time.", InvErrID (), iotaInv.Value ()))
	InvErr8 = func () (*err.Error) { return err.New ("The ID of a location was not provided.", InvErrID (), iotaInv.Value ()))
	InvErr9 = func () (*err.Error) { return err.New ("Request data validation failed.", InvErrID (), iotaInv.Value ()))

	OprErr0 = func () (*err.Error) { return err.New ("Database unreachable.", OprErrID (), iotaOpr.Value ())
	OprErr1 = func () (*err.Error) { return err.New ("Unable to check if all locations exist.", OprErrID (), iotaOpr.Value ())
	OprErr2 = func () (*err.Error) { return err.New ("Some locations do not exist.", OprErrID (), iotaOpr.Value ())
	OprErr3 = func () (*err.Error) { return err.New ("Unable to fetch the sensor IDs of all locations.", OprErrID (), iotaOpr.Value ())
	OprErr4 = func () (*err.Error) { return err.New ("Unable to fetch the sensor ID of a location.", OprErrID (), iotaOpr.Value ())
	OprErr5 = func () (*err.Error) { return err.New ("Unable to fetch the states of all locations.", OprErrID (), iotaOpr.Value ())
	OprErr6 = func () (*err.Error) { return err.New ("Unable to fetch a state data.", OprErrID (), iotaOpr.Value ())
	OprErr7 = func () (*err.Error) { return err.New ("Unable to make request records a squaket.", OprErrID (), iotaOpr.Value ())
	OprErr8 = func () (*err.Error) { return err.New ("Unable to group request records, based on sensor.", OprErrID (), iotaOpr.Value ())
	OprErr9 = func () (*err.Error) { return err.New ("Unable to make records of a sensor a squaket.", OprErrID (), iotaOpr.Value ())
	OprErr10 = func () (*err.Error) { return err.New ("Unable to group records of a sensor, based on day.", OprErrID (), iotaOpr.Value ())
	OprErr11 = func () (*err.Error) { return err.New ("A sensor record could not be marshalled.", OprErrID (), iotaOpr.Value ())
}
