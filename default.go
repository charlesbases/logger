package logger

var defaultEntry Logger

// SetDefault 配置默认的 Logger 接口的日志记录器实例
//
// 参数:
//   opts... - 可选的配置函数，用于定制日志记录器行为
//
// 示例:
//   SetDefault(
//       WithCallerSkip(2),
//       WithName("app"),
//       WithWriter(os.Stdout, fileWriter),
//       WithColorful(false),
//       WithMinLevel("debug"),
//   )
func SetDefault(opts ...Option) {
	options := []Option{WithCallerSkip(2)}
	defaultEntry = New(append(options, opts...)...)
}

// Debug logs a message at [zapcore.DebugLevel].
func Debug(v ...interface{}) {
	defaultEntry.Debug(v...)
}

// Debugf formats the message according to the format specifier
// and logs it at [zapcore.DebugLevel].
func Debugf(format string, v ...interface{}) {
	defaultEntry.Debugf(format, v...)
}

// Info logs a message at [zapcore.InfoLevel].
func Info(v ...interface{}) {
	defaultEntry.Info(v...)
}

// Infof formats the message according to the format specifier
// and logs it at [zapcore.InfoLevel].
func Infof(format string, v ...interface{}) {
	defaultEntry.Infof(format, v...)
}

// Warn logs a message at [zapcore.WarnLevel].
func Warn(v ...interface{}) {
	defaultEntry.Warn(v...)
}

// Warnf formats the message according to the format specifier
// and logs it at [zapcore.WarnLevel].
func Warnf(format string, v ...interface{}) {
	defaultEntry.Warnf(format, v...)
}

// Error logs a message at [zapcore.ErrorLevel].
func Error(v ...interface{}) {
	defaultEntry.Error(v...)
}

// Errorf formats the message according to the format specifier
// and logs it at [zapcore.ErrorLevel].
func Errorf(format string, v ...interface{}) {
	defaultEntry.Errorf(format, v...)
}

// Panic logs a message at [zapcore.PanicLevel] and panics.
func Panic(v ...interface{}) {
	defaultEntry.Panic(v...)
}

// Panicf formats the message according to the format specifier
// and panics.
func Panicf(format string, v ...interface{}) {
	defaultEntry.Panicf(format, v...)
}

// Fatal logs a message at [zapcore.FatalLevel] and calls os.Exit.
func Fatal(v ...interface{}) {
	defaultEntry.Fatal(v...)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func Fatalf(format string, v ...interface{}) {
	defaultEntry.Fatalf(format, v...)
}
