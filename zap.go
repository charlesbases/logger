package logger

import (
	"context"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger .
type Logger struct {
	pool *sync.Pool
	hook ContextHook

	sugared *zap.SugaredLogger
}

// wrap .
func wrap(v string) string {
	if len(v) != 0 {
		var b strings.Builder
		b.Grow(len(v) + 2)
		b.WriteString("[")
		b.WriteString(v)
		b.WriteString("]")
		return b.String()
	}
	return ""
}

// New .
func New(opts ...option) *Logger {
	var options = configuration(opts...)

	// 编码器
	encodercfg := zap.NewProductionEncoderConfig()
	encodercfg.EncodeTime = zapcore.TimeEncoderOfLayout(wrap(defaultDateFormat))
	encodercfg.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(shortName(level)(options.colourful))
	}
	encodercfg.EncodeCaller = zapcore.ShortCallerEncoder
	encodercfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(encodercfg)

	// 日志级别
	level := zap.LevelEnablerFunc(
		func(lv zapcore.Level) bool {
			return level(lv) >= options.minlevel
		},
	)

	// output-console
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)

	// output-writer
	if options.writer != nil {
		core = zapcore.NewTee([]zapcore.Core{core, zapcore.NewCore(encoder, zapcore.AddSync(options.writer), level)}...)
	}

	sugared := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(options.callerSkip)).Sugar()
	return &Logger{
		pool: &sync.Pool{
			New: func() interface{} {
				return sugared
			},
		},
		hook:    options.contextHook,
		sugared: sugared.Named(wrap(options.name)),
	}
}

// clone .
func (log *Logger) clone() *Logger {
	var copylog = *log
	return &copylog
}

// CallerSkip 添加调用层
func (log *Logger) CallerSkip(skip int) *Logger {
	if skip != 0 {
		l := log.clone()
		l.sugared = log.sugared.WithOptions(zap.AddCallerSkip(skip))
		return l
	}
	return log
}

// Context return new Logger with context
func (log *Logger) Context(ctx context.Context) *Logger {
	if log.hook != nil {
		return log.hook(ctx)(log)
	}
	return log
}

// Named 修改 name
// 注意：是修改，而不是 zap.Logger.Named() 的追加 name
func (log *Logger) Named(name string) *Logger {
	if len(name) != 0 {
		l := log.clone()
		base := l.pool.Get().(*zap.SugaredLogger)
		l.sugared = base.Named(wrap(name))
		l.pool.Put(base)
		return l
	}
	return log
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
