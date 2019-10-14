package main

import (
	"fmt"
	"gopkg.in/gorilla/mux.v1"
	"gopkg.in/qamarian-dtp/err.v0" //v0.4.0
	errLib "gopkg.in/qamarian-lib/err.v0" //v0.4.0
	"gopkg.in/qamarian-lib/str.v3"
	"net/http"
	"os"
	"time"
	"./lib" // pkg with InitAcct ()
	"./server"
)
func init () {
	errX := server.InitAcct ()
	if errX != nil {
		errY := err.New (`Package "./server" initialization failed.`, nil, nil, errX)
		str.PrintEtr (errLib.Fup (errY), "err", "init ()")
		os.Exit (1)
	}
}

func main () {
	str.PrintEtr ("Service has started up.", "std", "main ()")

	// | --
	readTimeout, errX := time.ParseDuration ("4m")
	readHeaderTimeout, errY := time.ParseDuration ("4m")
	writeTimeout, errZ := time.ParseDuration ("4m")
	idleTimeout, errA := time.ParseDuration ("4m")

	if errX != nil || errY != nil || errZ != nil || errA != nil {
		str.PrintEtr ("Buggy software. Ref: 0", "err", "main ()")
		os.Exit (1)
	}
	// -- |

	connAcceptor := &http.Server {
		ReadTimeout: readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout: idleTimeout,
		MaxHeaderBytes: 1 << 12,
	}

	router := mux.NewRouter ()
	router.HandleFunc ("/locHistory", server.ServiceRequestServer)

	connAcceptor.Handler = router

	conf, errM := lib.NewConf ()
	if errM != nil {
		errN := err.New ("Unable to load service configuration.", nil, nil, errM)
		str.PrintEtr (errLib.Fup (errN), "err", "main ()")
		os.Exit (1)
	}		

	connAcceptor.Addr = conf.Get ("service_addr") + ":" + conf.Get ("service_port")

	str.PrintEtr ("Service now running.", "std", "main ()")
	output := fmt.Sprintf ("Service addr: [%s]:%s", conf.Get ("service_addr"), conf.Get ("service_port"))
	str.PrintEtr (output, "nte", "main ()")

	errO := connAcceptor.ListenAndServeTLS (conf.Get ("service_tls_crt"), conf.Get ("service_tls_key"))
	if errO != nil && errO != http.ErrServerClosed {
		errP := err.New ("Shutdown due to an error.", nil, nil, errO)
		lib.StdErr.Write ([]byte (errLib.Fup (errP)))
		str.PrintEtr (errLib.Fup (errP), "err", "main ()")
		os.Exit (1)
	}

	str.PrintEtr ("Service has shutdown.", "nte", "main ()")
}
