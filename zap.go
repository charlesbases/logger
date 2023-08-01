package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger .
type logger struct {
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
func New(opts ...func(o *Options)) *logger {
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

	return &logger{
		skip: options.Skip,
		core: core,
		base: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.Skip+options.baseSkip)).Sugar().Named(warp(options.Name)),
	}
}

// Named .
func (log *logger) Named(name string, opts ...func(o *Options)) *logger {
	var options = new(Options)
	for _, opt := range opts {
		opt(options)
	}

	return &logger{
		skip: log.skip,
		core: log.core,
		base: zap.New(log.core, zap.AddCaller(), zap.AddCallerSkip(log.skip+options.Skip)).Sugar().Named(warp(name)),
	}
}

// Flush .
func (log *logger) Flush() {
	log.base.Sync()
}

// Trace .
func (log *logger) Trace(v ...interface{}) {
	log.base.Info(v...)
}

// Tracef .
func (log *logger) Tracef(format string, params ...interface{}) {
	log.base.Infof(format, params...)
}

// Debug .
func (log *logger) Debug(v ...interface{}) {
	log.base.Debug(v...)
}

// Debugf .
func (log *logger) Debugf(format string, params ...interface{}) {
	log.base.Debugf(format, params...)
}

// Info .
func (log *logger) Info(v ...interface{}) {
	log.base.Info(v...)
}

// Infof .
func (log *logger) Infof(format string, params ...interface{}) {
	log.base.Infof(format, params...)
}

// Warn .
func (log *logger) Warn(v ...interface{}) {
	log.base.Warn(v...)
}

// Warnf .
func (log *logger) Warnf(format string, params ...interface{}) {
	log.base.Warnf(format, params...)
}

// Error .
func (log *logger) Error(v ...interface{}) {
	log.base.Error(v...)
}

// Errorf .
func (log *logger) Errorf(format string, params ...interface{}) {
	log.base.Errorf(format, params...)
}

// Fatal .
func (log *logger) Fatal(v ...interface{}) {
	log.base.Fatal(v...)
}

// Fatalf .
func (log *logger) Fatalf(format string, params ...interface{}) {
	log.base.Fatalf(format, params...)
}
