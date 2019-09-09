package logger

import (
	"fmt"
	"log"
	"gopkg.in/qamarian-lib/str.v2"
	"gopkg.in/qamarian-mmp/rxlib.v0"
	"strconv"
	"time"
)

// Logger () is a Rexa main which record strings sent to it, as logs.
func Logger (key rxlib.Key) () {
	// Fetching a writer which logs can be forwarded to. { ...
	writer, errX := provideWriter ()
	if errX != nil {
		errMssg := fmt.Sprintf ("Unable to fetch a writer to which logs would " +
			"be forwarded. [%s]", errX.Error ())
		key.StartupFailed (errMssg)
		return
	}
	// ... }
	logger := log.New (writer, "", log.Ldate | log.Ltime)
	key.NowRunning ()
	// Recording all logs that are sent. { ...
	for {
		key.Wait ()
		mssg, errY := key.Read ()
		if errY != nil {
			errMssg := fmt.Sprintf ("[%s] Error occured while trying to " +
				"fetch a new message. [%s]", currentTimeString (),
				errY.Error ())
			str.PrintEtr (errMssg, "err", "qamarian-rxm/logger")
			continue
		}
		if mssg == nil {
			continue
		}
		log, okX := mssg.(string)
		if okX == false {
			errMssg := fmt.Sprintf ("[%s] A log sent to this main, is not " +
				"a string.", currentTimeString ())
			str.PrintEtr (errMssg, "err", "qamarian-rxm/logger")
			continue
		}
		logger.Println (log)
		if key.CheckForShutdown () == true && key.Check () == false {
			key.IndicateShutdown ()
			return
		}
	}
	// ... }
}

// currentTime () returns the current time, in the following format:
// "Sep 8, 2019; 20:56:07"
func currentTimeString () (string) {
	currentTime := time.Now ()

	month := currentTime.Month ().String ()
	month = month [0:3]

	day := ""
	if currentTime.Day () < 10 {
		day = "0" + strconv.Itoa (currentTime.Day ())
	} else {
		day = strconv.Itoa (currentTime.Day ())
	}

	hour := ""
	if currentTime.Hour () < 10 {
		hour = "0" + strconv.Itoa (currentTime.Hour ())
	} else {
		hour = strconv.Itoa (currentTime.Hour ())
	}

	min := ""
	if currentTime.Minute () < 10 {
		min = "0" + strconv.Itoa (currentTime.Minute ())
	} else {
		min = strconv.Itoa (currentTime.Minute ())
        }

	sec := ""
	if currentTime.Second () < 10 {
		sec = "0" + strconv.Itoa (currentTime.Second ())
	} else {
		sec = strconv.Itoa (currentTime.Second ())
        }

	return fmt.Sprintf ("%s %s, %d; %s:%s:%s", month, day, currentTime.Year (), hour,
		min, sec)
}
