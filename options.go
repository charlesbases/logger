package logger

import (
	"context"
	"io"
)

const (
	// defaultCallerSkip .
	defaultCallerSkip = 1
	// defaultDateFormat date format
	defaultDateFormat = "2006-01-02 15:04:05.000"
)

// ContextHook return new Logger with context
type ContextHook func(ctx context.Context) string

// options .
type options struct {
	// name 名称
	name string
	// callerSkip 跳过的调用者数量
	callerSkip int
	// colourful 日志级别多彩显示
	colourful bool
	// writer others output
	writer io.Writer
	// contextHook parse name in context. eg: TestContextHook
	contextHook ContextHook
	// minlevel convert MinLevel to level
	minlevel level
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
			o.colourful = true
		},
	)
}

// WithMinLevel .
func WithMinLevel(v string) option {
	return funcOption(
		func(o *options) {
			if minlevel, found := string2level[v]; found {
				o.minlevel = minlevel
			}
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
	var options = &options{callerSkip: defaultCallerSkip, minlevel: minlevel}
	for _, opt := range opts {
		opt.apply(options)
	}
	return options
}
