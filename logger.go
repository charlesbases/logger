package logger

var base *Logger

// baseSkip 因为 base 对 Logger 再封装了一层，所以 defaultCallerSkip + 1
var baseSkip = func(o *Options) { o.baseSkip = 1 }

func init() {
	base = New(baseSkip)
}

// SetDefault .
func SetDefault(opts ...func(o *Options)) {
	if base != nil {
		base.Flush()
	}
	base = New(append(opts, baseSkip)...)
}

// Flush .
func Flush() {
	base.Flush()
}

// Named .
func Named(name string, opts ...func(o *Options)) *Logger {
	return base.Named(name, opts...)
}

// Trace .
func Trace(v ...interface{}) {
	base.Trace(v...)
}

// Tracef .
func Tracef(format string, params ...interface{}) {
	base.Tracef(format, params...)
}

// Debug .
func Debug(v ...interface{}) {
	base.Debug(v...)
}

// Debugf .
func Debugf(format string, params ...interface{}) {
	base.Debugf(format, params...)
}

// Info .
func Info(v ...interface{}) {
	base.Info(v...)
}

// Infof .
func Infof(format string, params ...interface{}) {
	base.Infof(format, params...)
}

// Warn .
func Warn(v ...interface{}) {
	base.Warn(v...)
}

// Warnf .
func Warnf(format string, params ...interface{}) {
	base.Warnf(format, params...)
}

// Error .
func Error(v ...interface{}) {
	base.Error(v...)
}

// Errorf .
func Errorf(format string, params ...interface{}) {
	base.Errorf(format, params...)
}

// Fatal .
func Fatal(v ...interface{}) {
	base.Fatal(v...)
}

// Fatalf .
func Fatalf(format string, params ...interface{}) {
	base.Fatalf(format, params...)
}
