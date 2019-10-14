package lib

import (
	"gopkg.in/qamarian-dtp/err.v0" // v0.4.0
	viperLib "gopkg.in/qamarian-lib/viper.v0" // v0.1.0
	"gopkg.in/spf13/afero.v1"
	"net"
	"strconv"
)

func NewConf () (output *Conf, somrErr error) {
	conf, errX := viperLib.NewFileViper (confFileName, "yaml")
	if errX != nil {
		return nil, err.New ("Unable to load conf file.", nil, nil, errX)
	}

	output = &Conf {}


// Section X -- Server configuration processing.

	// Processing conf data 'service_addr'.  ..1.. {
	if ! conf.IsSet ("service_addr") {
		return nil, err.New ("Conf data 'service_addr': Data not set.", nil, nil)
	}
	(*output) ["service_addr"] = conf.GetString ("service_addr")
	ipAddrX := net.ParseIP ((*output) ["service_addr"])
	if ipAddrX == nil {
		return nil, err.New ("Conf data 'service_addr': Inalid IPv4 address.", nil, nil)
	}
	if ipAddrX.To4 () == nil {
		return nil, err.New ("Conf data 'service_addr': Inalid IPv4 address.", nil, nil)
	}
	// .. }

	// Processing conf data 'service_port'.  ..1.. {
	if ! conf.IsSet ("service_port") {
		return nil, err.New ("Conf data 'service_port': Data not set.", nil, nil)
	}
	(*output) ["service_port"] = conf.GetString ("service_port")
	portX, okX := strconv.Atoi ((*output) ["service_port"])
	if okX != nil || portX > 65535 {
		return nil, err.New ("Conf data 'service_port': Invalid port number.", nil, nil)
	}
	if portX == 0 {
		return nil, err.New ("Conf data 'service_port': Port 0 can not be used.", nil, nil)
	}
	// .. }

	// Processing conf data 'service_tls_key'.  ..1.. {	
	if ! conf.IsSet ("service_tls_key") || conf.GetString ("service_tls_key") == "" {
		return nil, err.New ("Conf data 'service_tls_key': Data not set.", nil, nil)
	}
	(*output) ["service_tls_key"] = conf.GetString ("service_tls_key")
	okR, errR := afero.Exists (afero.NewOsFs (), (*output) ["service_tls_key"])
	if errR != nil {
		return nil, err.New ("Conf data 'service_tls_key': Unable to confirm existence of file.", nil, nil)
	} else if okR == false {
		return nil, err.New ("Conf data 'service_tls_key': File not found.", nil, nil)
	}
	// .. }

	// Processing conf data 'service_tls_crt'.  ..1.. {	
	if ! conf.IsSet ("service_tls_crt") || conf.GetString ("service_tls_crt") == "" {
		return nil, err.New ("Conf data 'service_tls_crt': Data not set.", nil, nil)
	}
	(*output) ["service_tls_crt"] = conf.GetString ("service_tls_crt")
	okS, errS := afero.Exists (afero.NewOsFs (), (*output) ["service_tls_crt"])
	if errS != nil {
		return nil, err.New ("Conf data 'service_tls_crt': Unable to confirm existence of file.", nil, nil)
	} else if okS == false {
		return nil, err.New ("Conf data 'service_tls_crt': File not found.", nil, nil)
	}
	// .. }


// Section Y
	// Processing conf data 'dbms_pub_key'.  ..1.. {	
	if ! conf.IsSet ("dbms_pub_key") || conf.GetString ("dbms_pub_key") == "" {
		return nil, err.New ("Conf data 'dbms_pub_key': Data not set.", nil, nil)
	}
	(*output) ["dbms_pub_key"] = conf.GetString ("dbms_pub_key")
	okK, errK := afero.Exists (afero.NewOsFs (), (*output) ["dbms_pub_key"])
	if  errK != nil {
		return nil, err.New ("Conf data 'dbms_pub_key': Unable to confirm existence of file.", nil, nil)
	} else if okK == false {
		return nil, err.New ("Conf data 'dbms_pub_key': File not found.", nil, nil)
	}
	// .. }

	// Processing conf data 'dbms_addr'.  ..1.. {
	if ! conf.IsSet ("dbms_addr") {
		return nil, err.New ("Conf data 'dbms_addr': Data not set.", nil, nil)
	}
	(*output) ["dbms_addr"] = conf.GetString ("dbms_addr")
	ipAddrY := net.ParseIP ((*output) ["dbms_addr"])
	if ipAddrY == nil {
		return nil, err.New ("Conf data 'dbms_addr': Inalid IPv4/v6 address.", nil, nil)
	}
	// .. }

	// Processing conf data 'dbms_port'.  ..1.. {
	if ! conf.IsSet ("dbms_port") {
		return nil, err.New ("Conf data 'dbms_port': Data not set.", nil, nil)
	}
	(*output) ["dbms_port"] = conf.GetString ("dbms_port")
	portP, errP := strconv.Atoi ((*output) ["dbms_port"])
	if errP != nil || portP > 65535 {
		return nil, err.New ("Conf data 'dbms_port': Port number might be invalid.", nil, nil, errP)
	}
	// .. }

	// Processing conf data 'username'.  ..1.. {
	if ! conf.IsSet ("username") || conf.GetString ("username") == "" {
		return nil, err.New ("Conf data 'username': Data not set.", nil, nil)
	}
	(*output) ["username"] = conf.GetString ("username")
	// .. }

	// Processing conf data 'pass'.  ..1.. {
	if ! conf.IsSet ("pass") || conf.GetString ("pass") == "" {
		return nil, err.New ("Conf data 'pass': Data not set.", nil, nil)
	}
	(*output) ["pass"] = conf.GetString ("pass")
	// .. }

	// Processing conf data 'db'.  ..1.. {
	if ! conf.IsSet ("db") || conf.GetString ("db") == "" {
		return nil, err.New ("Conf data 'db': Data not set.", nil, nil)
	}
	(*output) ["pass"] = conf.GetString ("db")
	// .. }

	// Processing conf data 'conn_timeout'.  ..1.. {
	if ! conf.IsSet ("conn_timeout") {
		return nil, err.New ("Conf data 'conn_timeout': Data not set.", nil, nil)
	}
	(*output) ["conn_timeout"] = conf.GetString ("conn_timeout")
	timeoutA, errA := strconv.Atoi ((*output) ["conn_timeout"])
	if errA != nil {
		return nil, err.New ("Conf data 'conn_timeout': Value seems invalid.", nil, nil, errA)
	}
	if timeoutA == 0 {
		return nil, err.New ("Conf data 'conn_timeout': Timeout can not be zero.", nil, nil)
	}
	// .. }

	// Processing conf data 'write_timeout'.  ..1.. {
	if ! conf.IsSet ("write_timeout") {
		return nil, err.New ("Conf data 'write_timeout': Data not set.", nil, nil)
	}
	(*output) ["write_timeout"] = conf.GetString ("write_timeout")
	timeoutB, errB := strconv.Atoi ((*output) ["write_timeout"])
	if errB != nil {
		return nil, err.New ("Conf data 'write_timeout': Value seems invalid.", nil, nil, errB)
	}
	if timeoutB == 0 {
		return nil, err.New ("Conf data 'write_timeout': Timeout can not be zero.", nil, nil)
	}
	// .. }

	// Processing conf data 'read_timeout'.  ..1.. {
	if ! conf.IsSet ("read_timeout") {
		return nil, err.New ("Conf data 'read_timeout': Data not set.", nil, nil)
	}
	(*output) ["read_timeout"] = conf.GetString ("read_timeout")
	timeoutC, errC := strconv.Atoi ((*output) ["read_timeout"])
	if errC != nil {
		return nil, err.New ("Conf data 'read_timeout': Value seems invalid.", nil, nil, errC)
	}
	if timeoutC == 0 {
		return nil, err.New ("Conf data 'read_timeout': Timeout can not be zero.", nil, nil)
	}
	// .. }

	return output, nil
}

type Conf map[string]string

func (c *Conf) Get (name string) (string) {
	output, _ := (*c) [name]
	return output
}

var (
	confFileName string = "./conf.yml"
)
