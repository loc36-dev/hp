package lib

import (
	"gopkg.in/qamarian-dto/err.v0" // v0.4.0
	viperLib "gopkg.in/qamarian-lib/viper.v0" // v0.1.0
	"gopkg.in/spf13/afero.v1"
	"net"
	"strconv"
)

func NewConf () (output *Conf, err error) {
	conf, errX := viperLib.NewFileViper (confFileName, "yaml")
	if errX != nil {
		return nil, err.New ("Unable to load conf file.", nil, nil, errX)
	}

	output = &Conf {map[string]string {}}


// Section X -- Server configuration processing.

	// Processing conf data 'service_addr'.  ..1.. {
	if ! conf.Isset ("service_addr") {
		return nil, err.New ("Conf data 'service_addr': Data not set.", nil, nil)
	}
	output ["service_addr"] := conf.GetString ("service_addr")
	ipAddrX := net.ParseIP (output ["service_addr"])
	if ipAddrX == nil {
		return nil, err.New ("Conf data 'service_addr': Inalid IPv4 address.", nil, nil)
	}
	if ipAddrX.To4 () === nil {
		return nil, err.New ("Conf data 'service_addr': Inalid IPv4 address.", nil, nil)
	}
	// .. }

	// Processing conf data 'service_port'.  ..1.. {
	if ! conf.Isset ("service_port") {
		return nil, err.New ("Conf data 'service_port': Data not set.", nil, nil)
	}
	output ["service_port"] := conf.GetString ("service_port")
	portX, okX := strconv.Atoi (output ["service_port"])
	if okX == false || portX > 65535 {
		return nil, err.New ("Conf data 'service_port': Invalid port number.")
	}
	if portX == 0 {
		return nil, err.New ("Conf data 'service_port': Port 0 can not be used.")
	}
	// .. }

	// Processing conf data 'service_tls_key'.  ..1.. {	
	if ! conf.Isset ("service_tls_key") || conf.GetString ("service_tls_key") == "" {
		return nil, err.New ("Conf data 'service_tls_key': Data not set.", nil, nil)
	}
	output ["service_tls_key"] := conf.GetString ("service_tls_key")
	if ! afero.Exists (afero.NewOsFs (), output ["service_tls_key"])) {
		return nil, err.New ("Conf data 'service_tls_key': File not found.", nil, nil)
	}
	// .. }

	// Processing conf data 'service_tls_crt'.  ..1.. {	
	if ! conf.Isset ("service_tls_crt") || conf.GetString ("service_tls_crt") == "" {
		return nil, err.New ("Conf data 'service_tls_crt': Data not set.", nil, nil)
	}
	output ["service_tls_crt"] := conf.GetString ("service_tls_crt")
	if ! afero.Exists (afero.NewOsFs (), output ["service_tls_crt"])) {
		return nil, err.New ("Conf data 'service_tls_crt': File not found.", nil, nil)
	}
	// .. }


// Section Y
	// Processing conf data 'dbms_pub_key'.  ..1.. {	
	if ! conf.Isset ("dbms_pub_key") || conf.GetString ("dbms_pub_key") == "" {
		return nil, err.New ("Conf data 'dbms_pub_key': Data not set.", nil, nil)
	}
	output ["dbms_pub_key"] := conf.GetString ("dbms_pub_key")
	if ! afero.Exists (afero.NewOsFs (), output ["dbms_pub_key"])) {
		return nil, err.New ("Conf data 'dbms_pub_key': File not found.", nil, nil)
	}
	// .. }

	// Processing conf data 'dbms_addr'.  ..1.. {
	if ! conf.Isset ("dbms_addr") {
		return nil, err.New ("Conf data 'dbms_addr': Data not set.", nil, nil)
	}
	output ["dbms_addr"] := conf.GetString ("dbms_addr")
	ipAddrY := net.ParseIP (output ["dbms_addr"])
	if ipAddrY == nil {
		return nil, err.New ("Conf data 'dbms_addr': Inalid IPv4/v6 address.", nil, nil)
	}
	// .. }

	// Processing conf data 'dbms_port'.  ..1.. {
	if ! conf.Isset ("dbms_port") {
		return nil, err.New ("Conf data 'dbms_port': Data not set.", nil, nil)
	}
	output ["dbms_port"] := conf.GetString ("dbms_port")
	portX, okX := strconv.Atoi (output ["dbms_port"])
	if okX == false || portX > 65535 {
		return nil, err.New ("Conf data 'dbms_port': Invalid port number.")
	}
	// .. }

	// Processing conf data 'username'.  ..1.. {
	if ! conf.Isset ("username") || conf.GetString ("username") == "" {
		return nil, err.New ("Conf data 'username': Data not set.", nil, nil)
	}
	output ["username"] := conf.GetString ("username")
	// .. }

	// Processing conf data 'pass'.  ..1.. {
	if ! conf.Isset ("pass") || conf.GetString ("pass") == "" {
		return nil, err.New ("Conf data 'pass': Data not set.", nil, nil)
	}
	output ["pass"] := conf.GetString ("pass")
	// .. }

	// Processing conf data 'db'.  ..1.. {
	if ! conf.Isset ("db") || conf.GetString ("db") == "" {
		return nil, err.New ("Conf data 'db': Data not set.", nil, nil)
	}
	output ["pass"] := conf.GetString ("db")
	// .. }

	// Processing conf data 'conn_timeout'.  ..1.. {
	if ! conf.Isset ("conn_timeout") {
		return nil, err.New ("Conf data 'conn_timeout': Data not set.", nil, nil)
	}
	output ["conn_timeout"] := conf.GetString ("conn_timeout")
	timeoutA, okA := strconv.Atoi (output ["conn_timeout"])
	if okA == false {
		return nil, err.New ("Conf data 'conn_timeout': Invalid value.")
	}
	if timeoutA == 0 {
		return nil, err.New ("Conf data 'conn_timeout': Timeout can not be zero.")
	}
	// .. }

	// Processing conf data 'write_timeout'.  ..1.. {
	if ! conf.Isset ("write_timeout") {
		return nil, err.New ("Conf data 'write_timeout': Data not set.", nil, nil)
	}
	output ["write_timeout"] := conf.GetString ("write_timeout")
	timeoutB, okB := strconv.Atoi (output ["write_timeout"])
	if okB == false {
		return nil, err.New ("Conf data 'write_timeout': Invalid value.")
	}
	if timeoutB == 0 {
		return nil, err.New ("Conf data 'write_timeout': Timeout can not be zero.")
	}
	// .. }

	// Processing conf data 'read_timeout'.  ..1.. {
	if ! conf.Isset ("read_timeout") {
		return nil, err.New ("Conf data 'read_timeout': Data not set.", nil, nil)
	}
	output ["read_timeout"] := conf.GetString ("read_timeout")
	timeoutC, okC := strconv.Atoi (output ["read_timeout"])
	if okC == false {
		return nil, err.New ("Conf data 'read_timeout': Invalid value.")
	}
	if timeoutC == 0 {
		return nil, err.New ("Conf data 'read_timeout': Timeout can not be zero.")
	}
	// .. }

	return output, nil
}

type Conf struct {
	value map[string]string
}

func (c *Conf) Get (name string) (string) {
	output, _ := c.value [name]
	return output
}

var (
	confFileName string = "./confFile.yml"
)
