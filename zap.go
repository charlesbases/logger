package logger

import (
	"os"

	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	_ "go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	zopts  *zapoptions
	logger *zap.SugaredLogger
}

// zapoptions .
type zapoptions struct {
	skip    int
	console zapcore.Core
}

// warp .
func warp(name string) string {
	if len(name) != 0 {
		return "[" + name + "]"
	}
	return name
}

// New .
func New(opts ...Option) *Logger {
	var options = defaultOption()
	for _, opt := range opts {
		opt(options)
	}

	// 编码器
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout(warp(DefaultDateFormat))
	cfg.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(render[convertZapLevel(level)])
	}
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(cfg)

	// 日志级别
	level := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		var l = convertZapLevel(lv)
		return l >= options.minlevel && l <= options.maxlevel
	})

	var core zapcore.Core

	// output-console
	console := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	// output-writer
	if len(options.Writers) != 0 {
		cores := []zapcore.Core{console}
		for _, w := range options.Writers {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(w), level))
		}
		core = zapcore.NewTee(cores...)
	} else {
		core = console
	}

	return &Logger{
		zopts: &zapoptions{
			skip:    options.Skip,
			console: console,
		},
		logger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.Skip)).Sugar().Named(warp(options.Service)),
	}
}

// Name .
func (l *Logger) Name(v string) *Logger {
	return &Logger{
		zopts:  l.zopts,
		logger: zap.New(l.zopts.console, zap.AddCaller(), zap.AddCallerSkip(l.zopts.skip-1)).Sugar().Named(warp(v)),
	}
}

// Flush .
func (l *Logger) Flush() {
	l.logger.Sync()
}

// Trace .
func (l *Logger) Trace(v ...interface{}) {
	l.logger.Info(v...)
}

// Tracef .
func (l *Logger) Tracef(format string, params ...interface{}) {
	l.logger.Infof(format, params...)
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
