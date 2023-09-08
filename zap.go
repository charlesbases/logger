package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	base    *zap.Logger
	sugared *zap.SugaredLogger
}

// warp .
func warp(v string) string {
	if len(v) != 0 {
		var b strings.Builder
		b.Grow(len(v) + 2)
		b.WriteString("[")
		b.WriteString(v)
		b.WriteString("]")
		return b.String()
	}
	return v
}

// New .
func New(opts ...func(o *Options)) *Logger {
	var options = option(opts...)

	// 编码器
	encodercfg := zap.NewProductionEncoderConfig()
	encodercfg.EncodeTime = zapcore.TimeEncoderOfLayout(warp(defaultDateFormat))
	encodercfg.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(render[level])
	}
	encodercfg.EncodeCaller = zapcore.ShortCallerEncoder
	encodercfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(encodercfg)

	// 日志级别
	level := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return level(lv) >= options.minlevel
	})

	// output-console
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	// output-writer
	if options.Writer != nil {
		core = zapcore.NewTee([]zapcore.Core{core, zapcore.NewCore(encoder, zapcore.AddSync(options.Writer), level)}...)
	}

	base := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.Skip))
	sugared := base.Sugar()

	if len(options.Name) != 0 {
		sugared = sugared.Named(warp(options.Name))
	}

	return &Logger{base: base, sugared: sugared}
}

// Named 修改 name
// 注意：是修改，而不是 zap.Logger.Named() 的追加 name
func (log *Logger) Named(name string, opts ...func(o *Options)) *Logger {
	var options = new(Options)
	for _, opt := range opts {
		opt(options)
	}

	sugared := log.base.Sugar().Named(warp(name))
	if options.Skip != 0 {
		sugared = sugared.WithOptions(zap.AddCallerSkip(options.Skip))
	}
	return &Logger{base: log.base, sugared: sugared}
}

// CallerSkip 添加调用层
func (log *Logger) CallerSkip(skip int) *Logger {
	return &Logger{base: log.base, sugared: log.sugared.WithOptions(zap.AddCallerSkip(skip))}
}

// Flush .
func (log *Logger) Flush() {
	log.sugared.Sync()
}

// Debug .
func (log *Logger) Debug(v ...interface{}) {
	log.sugared.Debug(v...)
}

// Debugf .
func (log *Logger) Debugf(format string, params ...interface{}) {
	log.sugared.Debugf(format, params...)
}

// Info .
func (log *Logger) Info(v ...interface{}) {
	log.sugared.Info(v...)
}

// Infof .
func (log *Logger) Infof(format string, params ...interface{}) {
	log.sugared.Infof(format, params...)
}

// Warn .
func (log *Logger) Warn(v ...interface{}) {
	log.sugared.Warn(v...)
}

// Warnf .
func (log *Logger) Warnf(format string, params ...interface{}) {
	log.sugared.Warnf(format, params...)
}

// Error .
func (log *Logger) Error(v ...interface{}) {
	log.sugared.Error(v...)
}

// Errorf .
func (log *Logger) Errorf(format string, params ...interface{}) {
	log.sugared.Errorf(format, params...)
}

// Fatal .
func (log *Logger) Fatal(v ...interface{}) {
	log.sugared.Fatal(v...)
}

// Fatalf .
func (log *Logger) Fatalf(format string, params ...interface{}) {
	log.sugared.Fatalf(format, params...)
}
