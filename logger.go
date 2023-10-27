package logger

import "context"

var defaultLogger = New()

// SetDefault .
func SetDefault(opts ...func(o *Options)) {
	if defaultLogger != nil {
		defaultLogger.Flush()
	}
	defaultLogger = New(opts...)
}

// CallerSkip .
func CallerSkip(skip int) *Logger {
	return defaultLogger.CallerSkip(skip)
}

// WithContext .
func WithContext(ctx context.Context) *Logger {
	return defaultLogger.WithContext(ctx)
}

// Named .
func Named(name string) *Logger {
	return defaultLogger.Named(name)
}

// Flush .
func Flush() {
	defaultLogger.Flush()
}

// Debug .
func Debug(v ...interface{}) {
	defaultLogger.sugared.Debug(v...)
}

// Debugf .
func Debugf(format string, params ...interface{}) {
	defaultLogger.sugared.Debugf(format, params...)
}

// Info .
func Info(v ...interface{}) {
	defaultLogger.sugared.Info(v...)
}

// Infof .
func Infof(format string, params ...interface{}) {
	defaultLogger.sugared.Infof(format, params...)
}

// Warn .
func Warn(v ...interface{}) {
	defaultLogger.sugared.Warn(v...)
}

// Warnf .
func Warnf(format string, params ...interface{}) {
	defaultLogger.sugared.Warnf(format, params...)
}

// Error .
func Error(v ...interface{}) {
	defaultLogger.sugared.Error(v...)
}

// Errorf .
func Errorf(format string, params ...interface{}) {
	defaultLogger.sugared.Errorf(format, params...)
}

// Fatal .
func Fatal(v ...interface{}) {
	defaultLogger.sugared.Fatal(v...)
}

// Fatalf .
func Fatalf(format string, params ...interface{}) {
	defaultLogger.sugared.Fatalf(format, params...)
}
