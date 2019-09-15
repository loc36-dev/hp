package main

import (
	"gopkg.in/qamarian-mmp/rxlib.v0"
	// Import mains below ...
	"./logger"
	"./server"

)

var (
	mains []*rxlib.Register = []*rxlib.Register {
		rxlib.NewRegister ("logger", []string {}, logger.Logger),
		rxlib.NewRegister ("server", []string {"logger"}, server.Server),
	}
	osLog rxlib.RxLog = newLog ()
)
