package logger

import "context"

var logger *Logger

func init() {
	logger = New()
}

// SetDefault .
func SetDefault(opts ...func(o *Options)) {
	if logger != nil {
		logger.Flush()
	}
	logger = New(opts...)
}

// CallerSkip .
func CallerSkip(skip int) *Logger {
	return logger.CallerSkip(skip)
}

// WithContext .
func WithContext(ctx context.Context) *Logger {
	return logger.WithContext(ctx)
}

// Named .
func Named(name string, opts ...func(o *Options)) *Logger {
	return logger.Named(name, opts...)
}

// Flush .
func Flush() {
	logger.Flush()
}

// Debug .
func Debug(v ...interface{}) {
	logger.sugared.Debug(v...)
}

// Debugf .
func Debugf(format string, params ...interface{}) {
	logger.sugared.Debugf(format, params...)
}

// Info .
func Info(v ...interface{}) {
	logger.sugared.Info(v...)
}

// Infof .
func Infof(format string, params ...interface{}) {
	logger.sugared.Infof(format, params...)
}

// Warn .
func Warn(v ...interface{}) {
	logger.sugared.Warn(v...)
}

// Warnf .
func Warnf(format string, params ...interface{}) {
	logger.sugared.Warnf(format, params...)
}

// Error .
func Error(v ...interface{}) {
	logger.sugared.Error(v...)
}

// Errorf .
func Errorf(format string, params ...interface{}) {
	logger.sugared.Errorf(format, params...)
}

// Fatal .
func Fatal(v ...interface{}) {
	logger.sugared.Fatal(v...)
}

// Fatalf .
func Fatalf(format string, params ...interface{}) {
	logger.sugared.Fatalf(format, params...)
}
