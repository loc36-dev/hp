package logger

import (
	"fmt"
	"io"
)

// provideWriter () is called by the logger, to provide an io.Writer that logs should be
// forwarded to.
//
// Do not forget to provide an implementation for this function, when using it in your
// Rexa-based Software.
//
// On success, an io.Writer and a nil error should be returned. On failure, an io.Writer
// and an error should be returned.
func provideWriter () (io.Writer, error) {
	// Some implementation
	return testWriter {}, nil
}

// testWriter is just a data type placed here to faciliate the testing of this package.
// Delete this data type when coding your implementation of provideWriter ().
type testWriter struct {}

func (t testWriter) Write (data []byte) (int, error) {
	fmt.Print ("Package 'github.com/qamarian-rxm/logger' is being tested. Log: ")
	fmt.Print (string (data), "\n")
	return 1, nil
}
