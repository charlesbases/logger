package logger

import (
	"context"
	"io"

	"go.uber.org/zap/zapcore"
)

const (
	// defaultCallerSkip .
	defaultCallerSkip = 1
	// defaultDateFormat date format
	defaultDateFormat = "2006-01-02 15:04:05.000 Z07:00"
)

// ContextHook return new Logger with context
type ContextHook func(ctx context.Context) string

// options .
type options struct {
	// name 名称
	name string
	// callerSkip 跳过的调用者数量
	callerSkip int
	// levelEncoder .
	levelEncoder func(code zapcore.Level) string
	// writer others output
	writer io.Writer
	// contextHook parse name in context. eg: TestContextHook
	contextHook ContextHook
	// minlevel convert MinLevel to levelEncoder
	minlevel zapcore.Level
}

// option .
type option interface {
	apply(o *options)
}

// funcOption .
type funcOption func(o *options)

// apply .
func (f funcOption) apply(o *options) () {
	f(o)
}

// WithName .
func WithName(v string) option {
	return funcOption(
		func(o *options) {
			o.name = v
		},
	)
}

// WithCallerSkip .
func WithCallerSkip(v int) option {
	return funcOption(
		func(o *options) {
			o.callerSkip = v
		},
	)
}

// WithColorful .
func WithColorful() option {
	return funcOption(
		func(o *options) {
			o.levelEncoder = colorfulEncoder
		},
	)
}

// WithMinLevel .
func WithMinLevel(v string) option {
	return funcOption(
		func(o *options) {
			minlevel := zapcore.DebugLevel
			switch v {
			case "debug":
				minlevel = zapcore.DebugLevel
			case "info":
				minlevel = zapcore.InfoLevel
			case "warn":
				minlevel = zapcore.WarnLevel
			case "error":
				minlevel = zapcore.ErrorLevel
			case "fatal":
				minlevel = zapcore.FatalLevel
			}
			o.minlevel = minlevel
		},
	)
}

// WithWriter .
func WithWriter(v io.Writer) option {
	return funcOption(
		func(o *options) {
			o.writer = v
		},
	)
}

// WithContextHook .
func WithContextHook(hook ContextHook) option {
	return funcOption(
		func(o *options) {
			o.contextHook = hook
		},
	)
}

// configuration .
func configuration(opts ...option) *options {
	var options = &options{
		callerSkip:   defaultCallerSkip,
		levelEncoder: originEncoder,
		minlevel:     zapcore.DebugLevel,
	}
	for _, opt := range opts {
		opt.apply(options)
	}
	return options
}
