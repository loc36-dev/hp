package lib

import (
	"io"
	"log"
	"os"
)

var (
	StdErr io.Writer = os.Stderr
)

func init () {
	StdErr = log.New (StdErr, "", log.Ldate | log.Ltime).Writer ()
}
