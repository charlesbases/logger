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
func (l *Logger) configure(opts ...Option) {
	var options = defaultOption()
	for _, opt := range opts {
		opt(options)
	}
	l.opts = options

	// 编码器
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("[" + DefaultDateFormat + "]")
	cfg.EncodeLevel = l.color
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.ConsoleSeparator = " "
	encoder := zapcore.NewConsoleEncoder(cfg)

	// 日志级别
	level := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return true
	})

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),                                                          // console
			zapcore.NewCore(encoder, zapcore.AddSync(NewFileWriter(l.opts.FilePath, FileWriterWithTTL(l.opts.MaxRolls))), level), // file-writer
		),
		zap.AddCaller(),
		zap.AddCallerSkip(l.opts.Skip),
	)

	if len(l.opts.Service) != 0 {
		logger = logger.Named("[" + l.opts.Service + "]")
	}
	l.logger = logger.Sugar()
}

// color .
func (l *Logger) color(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var level level
	switch lv {
	case zapcore.DebugLevel:
		level = levelDBG
	case zapcore.InfoLevel:
		level = levelINF
	case zapcore.WarnLevel:
		level = levelWRN
	case zapcore.ErrorLevel:
		level = levelERR
	case zapcore.DPanicLevel:
		level = levelFAT
	case zapcore.PanicLevel:
		level = levelFAT
	case zapcore.FatalLevel:
		level = levelFAT
	default:
		level = levelTRC
	}

	enc.AppendString(level.short())
}

// Trace .
func (l *Logger) Trace(v ...interface{}) {
	// zap does not support trace
}

// Tracef .
func (l *Logger) Tracef(format string, params ...interface{}) {
	// zap does not support trace
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
