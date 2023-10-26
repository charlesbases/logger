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
type ContextHook func(ctx context.Context) func(l *Logger) *Logger

// Options .
type Options struct {
	// Name 名称
	Name string
	// CallerSkip 跳过的调用者数量
	CallerSkip int
	// MinLevel 允许的最小日志级别. default: "debug"
	MinLevel string
	// Colourful 日志级别多彩显示
	Colourful bool
	// Writer others output
	Writer io.Writer
	// ContextHook return new Logger with context. eg: TestContextHook
	ContextHook ContextHook
	// minlevel convert MinLevel to level
	minlevel level
}

// configuration .
func configuration(opts ...func(o *Options)) *Options {
	var options = &Options{CallerSkip: defaultCallerSkip, minlevel: minlevel}
	for _, opt := range opts {
		opt(options)
	}
	if len(options.MinLevel) != 0 {
		if minlevel, found := string2level[options.MinLevel]; found {
			options.minlevel = minlevel
		}
	}
	return options
}
