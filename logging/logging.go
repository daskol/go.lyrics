// Package logging provides common logging facility across the repo.
package logging

import "github.com/sirupsen/logrus"

// Default logger.
var logger Logger = New()

// Level is enumeration that encodes logging levels.
type Level int

const (
	None Level = iota
	Debug
	Info
	Warning
	Error
	Fatal
	Panic
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Warningf(format string, args ...interface{})

	SetLevel(level Level)
}

// Default returns common logger for all packages.
func Default() Logger {
	return logger
}

// New constructs new logger.
func New() Logger {
	logger := newSirupsenLogger()
	logger.SetLevel(Info)
	return logger
}

// sirupsenLogger is a wrapper over logrus.Logger in order to provide common
// routines like setting logging level.
type sirupsenLogger struct {
	*logrus.Logger
}

func newSirupsenLogger() *sirupsenLogger {
	var logger = new(sirupsenLogger)
	logger.Logger = logrus.New()
	return logger
}

func (s *sirupsenLogger) SetLevel(level Level) {
	switch level {
	case Debug:
		s.Logger.SetLevel(logrus.DebugLevel)
	case Info:
		s.Logger.SetLevel(logrus.InfoLevel)
	case Warning:
		s.Logger.SetLevel(logrus.WarnLevel)
	case Error:
		s.Logger.SetLevel(logrus.ErrorLevel)
	case Fatal:
		s.Logger.SetLevel(logrus.FatalLevel)
	case Panic:
		s.Logger.SetLevel(logrus.PanicLevel)
	}
}
