package logger

// Logger logget interface
type Logger interface {
	// Debug logs a message at [zapcore.DebugLevel].
	Debug(v ...interface{})

	// Debugf formats the message according to the format specifier
	// and logs it at [zapcore.DebugLevel].
	Debugf(format string, v ...interface{})

	// Info logs a message at [zapcore.InfoLevel].
	Info(v ...interface{})

	// Infof formats the message according to the format specifier
	// and logs it at [zapcore.InfoLevel].
	Infof(format string, v ...interface{})

	// Warn logs a message at [zapcore.WarnLevel].
	Warn(v ...interface{})

	// Warnf formats the message according to the format specifier
	// and logs it at [zapcore.WarnLevel].
	Warnf(format string, v ...interface{})

	// Error logs a message at [zapcore.ErrorLevel].
	Error(v ...interface{})

	// Errorf formats the message according to the format specifier
	// and logs it at [zapcore.ErrorLevel].
	Errorf(format string, v ...interface{})

	// Panic logs a message at [zapcore.PanicLevel] and panics.
	Panic(v ...interface{})

	// Panicf formats the message according to the format specifier
	// and panics.
	Panicf(format string, v ...interface{})

	// Fatal logs a message at [zapcore.FatalLevel] and calls os.Exit.
	Fatal(v ...interface{})

	// Fatalf formats the message according to the format specifier
	// and calls os.Exit.
	Fatalf(format string, v ...interface{})
}
