package logger

import "context"

var base *Logger

func init() {
	base = New()
}

// SetDefault .
func SetDefault(opts ...func(o *Options)) {
	if base != nil {
		base.Flush()
	}
	base = New(opts...)
}

// CallerSkip .
func CallerSkip(skip int) *Logger {
	return base.CallerSkip(skip)
}

// WithContext .
func WithContext(ctx context.Context) *Logger {
	return base.WithContext(ctx)
}

// Named .
func Named(name string) *Logger {
	return base.Named(name)
}

// Flush .
func Flush() {
	base.Flush()
}

// Debug .
func Debug(v ...interface{}) {
	base.sugared.Debug(v...)
}

// Debugf .
func Debugf(format string, params ...interface{}) {
	base.sugared.Debugf(format, params...)
}

// Info .
func Info(v ...interface{}) {
	base.sugared.Info(v...)
}

// Infof .
func Infof(format string, params ...interface{}) {
	base.sugared.Infof(format, params...)
}

// Warn .
func Warn(v ...interface{}) {
	base.sugared.Warn(v...)
}

// Warnf .
func Warnf(format string, params ...interface{}) {
	base.sugared.Warnf(format, params...)
}

// Error .
func Error(v ...interface{}) {
	base.sugared.Error(v...)
}

// Errorf .
func Errorf(format string, params ...interface{}) {
	base.sugared.Errorf(format, params...)
}

// Fatal .
func Fatal(v ...interface{}) {
	base.sugared.Fatal(v...)
}

// Fatalf .
func Fatalf(format string, params ...interface{}) {
	base.sugared.Fatalf(format, params...)
}
