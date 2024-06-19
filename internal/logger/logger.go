package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Init initializes the logger with a JSON formatter, stdout output, and info level.
//
// No parameters.
// No return values.
func Init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// Log returns the logger instance.
//
// No parameters.
// Returns a pointer to a logrus.Logger.
func Log() *logrus.Logger {
	return log
}
