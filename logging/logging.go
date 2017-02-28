package logging

// Logging is the logging interface providers need to implement.
type Logging interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}
