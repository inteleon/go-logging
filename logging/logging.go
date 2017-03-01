package logging

import (
	"io"
)

const (
	// DebugLogLevel is used to set the log level to debug.
	DebugLogLevel = "debug"

	// InfoLogLevel is used to set the log level to info.
	InfoLogLevel = "info"

	// WarningLogLevel is used to set the log level to warning.
	WarningLogLevel = "warning"

	// ErrorLogLevel is used to set the log level to error.
	ErrorLogLevel = "error"

	// FatalLogLevel is used to set the log level to fatal.
	FatalLogLevel = "fatal"

	// JSONFormatter is used to set the log formatting to JSON.
	JSONFormatter = "json"

	// TextFormatter is used to set the log formatting to text.
	TextFormatter = "text"
)

// Logging is the logging interface providers need to implement.
type Logging interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})

	SetOutput(io.Writer)
	SetLogLevel(string)
	SetFormatter(string)
}
