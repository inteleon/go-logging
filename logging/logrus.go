package logging

import (
	log "github.com/sirupsen/logrus"
	"io"
	"reflect"
)

// LogrusLogging is an easily testable logrus logging implementation.
type LogrusLogging struct {
	Log *log.Logger
}

// NewLogrusLogging initiates and returns a new logrus logging object.
func NewLogrusLogging(logLevel log.Level, output io.Writer, formatter log.Formatter) Logging {
	l := log.New()
	l.Level = logLevel
	l.Out = output

	if formatter != nil {
		l.Formatter = formatter
	}

	return &LogrusLogging{
		Log: l,
	}
}

// Debug prints out a debug message.
func (l *LogrusLogging) Debug(message string, params interface{}) {
	l.setup(params).Debug(message)
}

// Info prints out an info message.
func (l *LogrusLogging) Info(message string, params interface{}) {
	l.setup(params).Info(message)
}

// Warn prints out a warning message.
func (l *LogrusLogging) Warn(message string, params interface{}) {
	l.setup(params).Warn(message)
}

// Error prints out an error message.
func (l *LogrusLogging) Error(message string, params interface{}) {
	l.setup(params).Error(message)
}

// Fatal prints out a fatal message and then exits with exit code 1.
func (l *LogrusLogging) Fatal(message string, params interface{}) {
	l.setup(params).Fatal(message)
}

// setup is a helper function for sorting out the provided logging parameters.
func (l *LogrusLogging) setup(params interface{}) *log.Entry {
	if params != nil {
		switch reflect.TypeOf(params).Kind() {
		case reflect.Map:
			p := reflect.ValueOf(params)

			fields := map[string]interface{}{}

			for _, k := range p.MapKeys() {
				fields[k.String()] = p.MapIndex(k).Interface()
			}

			return l.Log.WithFields(fields)
		}
	}

	return log.NewEntry(l.Log)
}
