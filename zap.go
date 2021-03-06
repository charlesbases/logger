package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	opts   *Options
	logger *zap.SugaredLogger
}

// New .
func New(opts ...Option) *Logger {
	l := new(Logger)
	l.configure(opts...)
	return l
}

// configure .
func (l *Logger) configure(options ...Option) {
	var opts = defaultOption()
	for _, opt := range options {
		opt(opts)
	}
	l.opts = opts

	// 编码器
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("[" + DefaultDateFormat + "]")
	cfg.EncodeLevel = l.color
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(cfg)

	// 日志级别
	level := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		var ll = convertZapLevel(lv)
		return ll >= l.opts.minlevel && ll <= l.opts.maxlevel
	})

	cores := make([]zapcore.Core, 0, len(opts.Writers)+1)
	// console
	cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level))
	// others-writer
	for _, w := range opts.Writers {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), level))
	}

	logger := zap.New(
		zapcore.NewTee(cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(l.opts.Skip),
	)

	if len(l.opts.Service) != 0 {
		logger = logger.Named(fmt.Sprintf("[%s]", l.opts.Service))
	}
	l.logger = logger.Sugar()
}

// color .
func (l *Logger) color(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(render[convertZapLevel(lv)])
}

// Trace .
func (l *Logger) Trace(v ...interface{}) {
	// l.logger.Info(v...)
}

// Tracef .
func (l *Logger) Tracef(format string, params ...interface{}) {
	// l.logger.Infof(format, params...)
}

// Debug .
func (l *Logger) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

// Debugf .
func (l *Logger) Debugf(format string, params ...interface{}) {
	l.logger.Debugf(format, params...)
}

// Info .
func (l *Logger) Info(v ...interface{}) {
	l.logger.Info(v...)
}

// Infof .
func (l *Logger) Infof(format string, params ...interface{}) {
	l.logger.Infof(format, params...)
}

// Warn .
func (l *Logger) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

// Warnf .
func (l *Logger) Warnf(format string, params ...interface{}) {
	l.logger.Warnf(format, params...)
}

// Error .
func (l *Logger) Error(v ...interface{}) {
	l.logger.Error(v...)
}

// Errorf .
func (l *Logger) Errorf(format string, params ...interface{}) {
	l.logger.Errorf(format, params...)
}

// Fatal .
func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

// Fatalf .
func (l *Logger) Fatalf(format string, params ...interface{}) {
	l.logger.Fatalf(format, params...)
}
