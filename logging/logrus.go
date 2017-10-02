package logging

import (
	"io"
	"log/syslog"
	"reflect"

	log "github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

// LogrusLogging is an easily testable logrus logging implementation.
type LogrusLogging struct {
	Log     *log.Logger
	Options LogrusLoggingOptions
}

// LogrusLoggingOptions is used for configuring the Logrus logging driver.
type LogrusLoggingOptions struct {
	Syslog *LogrusLoggingSyslogOptions
}

// LogrusLoggingSyslogOptions is used for configuring Logrus driver's syslog hook.
type LogrusLoggingSyslogOptions struct {
	Protocol string
	Address  string
	Priority syslog.Priority
	Tag      string
}

// NewLogrusLogging initiates and returns a new logrus logging object.
func NewLogrusLogging(options LogrusLoggingOptions) (Logging, error) {
	l := &LogrusLogging{
		Log:     log.New(),
		Options: options,
	}

	err := l.syslog()

	return l, err
}

func (l *LogrusLogging) syslog() error {
	if l.Options.Syslog == nil {
		return nil
	}

	hook, err := lSyslog.NewSyslogHook(
		l.Options.Syslog.Protocol,
		l.Options.Syslog.Address,
		l.Options.Syslog.Priority,
		l.Options.Syslog.Tag,
	)

	if err != nil {
		return err
	}

	l.Log.Hooks.Add(hook)

	return nil
}

// SetOutput sets the output of the logger - where to write to.
func (l *LogrusLogging) SetOutput(output io.Writer) {
	l.Log.Out = output
}

// SetLogLevel sets the logger log level.
func (l *LogrusLogging) SetLogLevel(logLevel string) {
	switch logLevel {
	case DebugLogLevel:
		l.Log.Level = log.DebugLevel

		break
	case InfoLogLevel:
		l.Log.Level = log.InfoLevel

		break
	case WarningLogLevel:
		l.Log.Level = log.WarnLevel

		break
	case ErrorLogLevel:
		l.Log.Level = log.ErrorLevel

		break
	case FatalLogLevel:
		l.Log.Level = log.FatalLevel

		break
	default:
		l.Log.Level = log.InfoLevel

		break
	}
}

// SetFormatter sets the logger formatting.
func (l *LogrusLogging) SetFormatter(formatter string) {
	switch formatter {
	case JSONFormatter:
		l.Log.Formatter = &log.JSONFormatter{}

		break
	case TextFormatter:
		l.Log.Formatter = &log.TextFormatter{}

		break
	default:
		l.Log.Formatter = &log.JSONFormatter{}

		break
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
