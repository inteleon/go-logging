package logging

// Logging is the logging interface providers need to implement.
type Logging interface {
	Debug(string, interface{})
	Info(string, interface{})
	Warn(string, interface{})
	Error(string, interface{})
	Fatal(string, interface{})
}
