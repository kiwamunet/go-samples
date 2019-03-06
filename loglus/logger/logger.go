package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func LogNew() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	// Log as JSON instead of the default ASCII formatter.
	l.SetFormatter(&logrus.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	l.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	l.SetLevel(logrus.DebugLevel)
	return l
}
