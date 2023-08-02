package logger

var logger *Logger

// baseSkip 因为 logger 对 Logger 再封装了一层，所以 defaultCallerSkip + 1
var baseSkip = func(o *Options) { o.baseSkip = 1 }

func init() {
	logger = New(baseSkip)
}

// SetDefault .
func SetDefault(opts ...func(o *Options)) {
	if logger != nil {
		logger.Flush()
	}
	logger = New(append(opts, baseSkip)...)
}

// Named .
func Named(name string, opts ...func(o *Options)) *Logger {
	return logger.Named(name, opts...)
}

// CallerSkip .
func CallerSkip(skip int) *Logger {
	return logger.CallerSkip(skip - 1)
}

// Flush .
func Flush() {
	logger.Flush()
}

// Trace .
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Tracef .
func Tracef(format string, params ...interface{}) {
	logger.Tracef(format, params...)
}

// Debug .
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf .
func Debugf(format string, params ...interface{}) {
	logger.Debugf(format, params...)
}

// Info .
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof .
func Infof(format string, params ...interface{}) {
	logger.Infof(format, params...)
}

// Warn .
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf .
func Warnf(format string, params ...interface{}) {
	logger.Warnf(format, params...)
}

// Error .
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf .
func Errorf(format string, params ...interface{}) {
	logger.Errorf(format, params...)
}

// Fatal .
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf .
func Fatalf(format string, params ...interface{}) {
	logger.Fatalf(format, params...)
}
