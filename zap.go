package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	skip int
	core zapcore.Core
	base *zap.SugaredLogger
}

// warp .
func warp(name string) string {
	if len(name) != 0 {
		return "[" + name + "]"
	}
	return name
}

// New .
func New(opts ...func(o *Options)) *Logger {
	var options = defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	if len(options.MaxLevel) != 0 {
		options.maxlevel = convertString(options.MaxLevel)
	}
	if len(options.MinLevel) != 0 {
		options.minlevel = convertString(options.MinLevel)
	}

	// 编码器
	encodercfg := zap.NewProductionEncoderConfig()
	encodercfg.EncodeTime = zapcore.TimeEncoderOfLayout(warp(defaultDateFormat))
	encodercfg.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(render[convertZapLevel(level)])
	}
	encodercfg.EncodeCaller = zapcore.ShortCallerEncoder
	encodercfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(encodercfg)

	// 日志级别
	level := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		var l = convertZapLevel(lv)
		return l >= options.minlevel && l <= options.maxlevel
	})

	// output-console
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	// output-writer
	if options.Writer != nil {
		core = zapcore.NewTee([]zapcore.Core{core, zapcore.NewCore(encoder, zapcore.AddSync(options.Writer), level)}...)
	}

	return &Logger{
		skip: options.Skip,
		core: core,
		base: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.Skip+options.baseSkip)).Sugar().Named(warp(options.Name)),
	}
}

// Named .
func (log *Logger) Named(name string, opts ...func(o *Options)) *Logger {
	var options = new(Options)
	for _, opt := range opts {
		opt(options)
	}

	return &Logger{
		skip: log.skip,
		core: log.core,
		base: zap.New(log.core, zap.AddCaller(), zap.AddCallerSkip(log.skip+options.Skip)).Sugar().Named(warp(name)),
	}
}

// Flush .
func (log *Logger) Flush() {
	log.base.Sync()
}

// Trace .
func (log *Logger) Trace(v ...interface{}) {
	log.base.Info(v...)
}

// Tracef .
func (log *Logger) Tracef(format string, params ...interface{}) {
	log.base.Infof(format, params...)
}

// Debug .
func (log *Logger) Debug(v ...interface{}) {
	log.base.Debug(v...)
}

// Debugf .
func (log *Logger) Debugf(format string, params ...interface{}) {
	log.base.Debugf(format, params...)
}

// Info .
func (log *Logger) Info(v ...interface{}) {
	log.base.Info(v...)
}

// Infof .
func (log *Logger) Infof(format string, params ...interface{}) {
	log.base.Infof(format, params...)
}

// Warn .
func (log *Logger) Warn(v ...interface{}) {
	log.base.Warn(v...)
}

// Warnf .
func (log *Logger) Warnf(format string, params ...interface{}) {
	log.base.Warnf(format, params...)
}

// Error .
func (log *Logger) Error(v ...interface{}) {
	log.base.Error(v...)
}

// Errorf .
func (log *Logger) Errorf(format string, params ...interface{}) {
	log.base.Errorf(format, params...)
}

// Fatal .
func (log *Logger) Fatal(v ...interface{}) {
	log.base.Fatal(v...)
}

// Fatalf .
func (log *Logger) Fatalf(format string, params ...interface{}) {
	log.base.Fatalf(format, params...)
}
