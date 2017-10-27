package logging

// Generates a logger for a specific Logrus log level
func (l *LogrusLogging) Logger(level string) Logger {
	logger := struct{ LogrusLogger }{}

	switch level {
	case DebugLogLevel:
		logger.printf = l.Log.Debugf
	case InfoLogLevel:
		logger.printf = l.Log.Infof
	case WarningLogLevel:
		logger.printf = l.Log.Warningf
	case ErrorLogLevel:
		logger.printf = l.Log.Errorf
	case FatalLogLevel:
		logger.printf = l.Log.Fatalf
	}

	return logger
}

type LogrusLogger struct {
	printf func(string, ...interface{})
}

func (l LogrusLogger) Printf(str string, args ...interface{}) {
	l.printf(str, args...)
}
