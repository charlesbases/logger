package logger

import (
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*zapLogger)(nil)

type zapLogger struct {
	name       string
	colorful   bool
	callerSkip int

	minLevel zapcore.Level

	// multiple WriteSyncer and sync calls
	writers []io.Writer

	sugared *zap.SugaredLogger
}

// New 创建并返回一个实现了 Logger 接口的日志记录器实例
//
// 参数:
//   opts... - 可选的配置函数，用于定制日志记录器行为
//
// 返回值:
//   Logger - 符合 Logger 接口的日志记录器
//
// 示例:
//   logger := New(
//       WithName("myapp"),
//       WithWriter(os.Stdout, fileWriter),
//       WithColorful(false),
//       WithMinLevel("debug"),
//   )
func New(opts ...Option) Logger {
	l := newSuger(opts...)
	core, err := l.initCore()
	if err != nil {
		log.Fatalln("init logger failed. ", err)
	}
	l.sugared = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(l.callerSkip)).Named(l.name).Sugar()
	return l
}

func newSuger(opts ...Option) *zapLogger {
	l := &zapLogger{
		callerSkip: 1,
		minLevel:   zapcore.DebugLevel,
		writers:    []io.Writer{os.Stdout},
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *zapLogger) initCore() (zapcore.Core, error) {
	config := zapcore.EncoderConfig{
		TimeKey:          "location",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		ConsoleSeparator: " ",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000 Z0700"),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
	}

	if l.colorful {
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	writers := make([]zapcore.WriteSyncer, 0, len(l.writers))
	for _, w := range l.writers {
		writers = append(writers, zapcore.AddSync(w))
	}

	return zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.NewMultiWriteSyncer(writers...), l.minLevel), nil
}

// Debug logs a message at [zapcore.DebugLevel].
func (l *zapLogger) Debug(v ...interface{}) {
	l.sugared.Debugln(v...)
}

// Debugf formats the message according to the format specifier
// and logs it at [zapcore.DebugLevel].
func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.sugared.Debugf(format, v...)
}

// Info logs a message at [zapcore.InfoLevel].
func (l *zapLogger) Info(v ...interface{}) {
	l.sugared.Infoln(v...)
}

// Infof formats the message according to the format specifier
// and logs it at [zapcore.InfoLevel].
func (l *zapLogger) Infof(format string, v ...interface{}) {
	l.sugared.Infof(format, v...)
}

// Warn logs a message at [zapcore.WarnLevel].
func (l *zapLogger) Warn(v ...interface{}) {
	l.sugared.Warnln(v...)
}

// Warnf formats the message according to the format specifier
// and logs it at [zapcore.WarnLevel].
func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.sugared.Warnf(format, v...)
}

// Error logs a message at [zapcore.ErrorLevel].
func (l *zapLogger) Error(v ...interface{}) {
	l.sugared.Errorln(v...)
}

// Errorf formats the message according to the format specifier
// and logs it at [zapcore.ErrorLevel].
func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.sugared.Errorf(format, v...)
}

// Panic logs a message at [zapcore.PanicLevel] and panics.
func (l *zapLogger) Panic(v ...interface{}) {
	l.sugared.Panicln(v...)
}

// Panicf formats the message according to the format specifier
// and panics.
func (l *zapLogger) Panicf(format string, v ...interface{}) {
	l.sugared.Panicf(format, v...)
}

// Fatal logs a message at [zapcore.FatalLevel] and calls os.Exit.
func (l *zapLogger) Fatal(v ...interface{}) {
	l.sugared.Fatalln(v...)
}

// Fatalf formats the message according to the format specifier
// and calls os.Exit.
func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.sugared.Fatalf(format, v...)
}
