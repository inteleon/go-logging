package logging

import (
	log "github.com/sirupsen/logrus"
	"io"
	"reflect"
)

// DebugLevel is used to set the logger to the debug log level.
const DebugLevel = log.DebugLevel

// InfoLevel is used to set the logger to the info log level.
const InfoLevel = log.InfoLevel

// WarnLevel is used to set the logger to the warn log level.
const WarnLevel = log.WarnLevel

// ErrorLevel is used to set the logger to the error log level.
const ErrorLevel = log.ErrorLevel

// FatalLevel is used to set the logger to the fatal log level.
const FatalLevel = log.FatalLevel

// LogrusLogging is an easily testable logrus logging implementation.
type LogrusLogging struct {
	Log *log.Logger
}

type LogrusLoggingOptions struct {
	LogLevel  log.Level
	Output    io.Writer
	Formatter log.Formatter
}

// NewLogrusLogging initiates and returns a new logrus logging object.
func NewLogrusLogging(options LogrusLoggingOptions) Logging {
	l := log.New()

	if options.LogLevel != 0 {
		l.Level = options.LogLevel
	}

	if options.Output != nil {
		l.Out = options.Output
	}

	if options.Formatter != nil {
		l.Formatter = options.Formatter
	}

	return &LogrusLogging{
		Log: l,
	}
}

// Debug prints out a debug message.
func (l *LogrusLogging) Debug(args ...interface{}) {
	if len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.String {
		l.setup(args[1:]).Debug(args[0].(string))
	}
}

// Info prints out an info message.
func (l *LogrusLogging) Info(args ...interface{}) {
	if len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.String {
		l.setup(args[1:]).Info(args[0].(string))
	}
}

// Warn prints out a warning message.
func (l *LogrusLogging) Warn(args ...interface{}) {
	if len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.String {
		l.setup(args[1:]).Warn(args[0].(string))
	}
}

// Error prints out an error message.
func (l *LogrusLogging) Error(args ...interface{}) {
	if len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.String {
		l.setup(args[1:]).Error(args[0].(string))
	}
}

// Fatal prints out a fatal message and then exits with exit code 1.
func (l *LogrusLogging) Fatal(args ...interface{}) {
	if len(args) > 0 && reflect.TypeOf(args[0]).Kind() == reflect.String {
		l.setup(args[1:]).Fatal(args[0].(string))
	}
}

// setup is a helper function for sorting out the provided logging parameters.
func (l *LogrusLogging) setup(params []interface{}) *log.Entry {
	if len(params) == 1 && params[0] != nil {
		switch reflect.TypeOf(params[0]).Kind() {
		case reflect.Map:
			p := reflect.ValueOf(params[0])

			fields := map[string]interface{}{}

			for _, k := range p.MapKeys() {
				fields[k.String()] = p.MapIndex(k).Interface()
			}

			return l.Log.WithFields(fields)
		}
	}

	return log.NewEntry(l.Log)
}
