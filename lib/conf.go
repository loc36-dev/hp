package lib

import (
	"gopkg.in/qamarian-dto/err.v0" // v0.4.0
	viperLib "gopkg.in/qamarian-lib/viper.v0" // v0.1.0
	"strconv"
)

func NewConf () (output *Conf, err error) {
	output = &Conf {map[string]string {}}

	// ..1.. {
	conf, errX := viperLib.NewFileViper (confFileName, "yaml")
	if errX != nil {
		return nil, err.New ("Unable to load conf file.", nil, nil, errX)
	}
	// ..1.. }

	// ..1.. {
	if ! conf.Isset ("conn_timeout") {
		return nil, err.New ("DB connection timeout is not set.", nil, nil)
	}
	data0 := conf.GetInt ()
	output.value ["conn_timeout"] = strconv.Itoa (data)
	// ..1.. }

	// ..1.. {
	if ! conf.Isset ("server_pub_key") {
		return nil, err.New ("Path of server public key not set.", nil, nil)
	}
	data1 := conf.GetString ()
	output.value ["server_pub_key"] = data1
	// ..1.. }

	// ..1.. {
	if ! conf.Isset ("db_write_timeout") {
		return nil, err.New ("Db write timeout is not set.", nil, nil)
	}
	data2 := conf.GetInt ()
	output.value ["db_write_timeout"] = strconv.Itoa (data2)
	// ..1.. }

	// ..1.. {
	if ! conf.Isset ("db_read_timeout") {
		return nil, err.New ("DB read timeout is not set.", nil, nil)
	}
	data3 := conf.GetInt ()
	output.value ["db_read_timeout"] = strconv.Itoa (data3)
	// ..1.. {
}

type Conf struct {
	value map[string]string
}

func (c *Conf) Get (name string) (string, error) {}



var (
	confFileName string = "./confFile.yml"
)
