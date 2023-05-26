package logger

import (
	"io"
)

const (
	// DefaultSkip .
	DefaultSkip = 1
	// DefaultDateFormat date format
	DefaultDateFormat = "2006-01-02 15:04:05.000"
)

// Options .
type Options struct {
	// Service 服务名
	Service string
	// Writers others output
	Writers []io.Writer
	// MinLevel 允许的最小日志级别. default: "Trace"
	MinLevel string
	// MaxLevel 允许的最大日志级别. default: "Fatal"
	MaxLevel string
	// Skip 跳过的调用者数量. default: DefaultSkip
	Skip int

	// 是否写入文件
	store bool
	// minlevel convert from MinLevel. default: _minlevel
	minlevel Level
	// maxlevel convert from MaxLevel. default: _maxlevel
	maxlevel Level
}

// defaultOption .
func defaultOption() *Options {
	return &Options{
		Skip: DefaultSkip,

		store:    false,
		minlevel: minlevel,
		maxlevel: maxlevel,
	}
}

type Option func(o *Options)

// WithSkip .
func WithSkip(skip int) Option {
	return func(o *Options) {
		o.Skip = skip
	}
}

// WithWriter .
func WithWriter(w ...io.Writer) Option {
	return func(o *Options) {
		if len(o.Writers) != 0 {
			o.Writers = append(o.Writers, w...)
		} else {
			o.Writers = w
		}
	}
}

// WithService .
func WithService(service string) Option {
	return func(o *Options) {
		o.Service = service
	}
}

// WithMinLevel .
func WithMinLevel(l Level) Option {
	return func(o *Options) {
		o.minlevel = l
	}
}

// WithMaxLevel .
func WithMaxLevel(l Level) Option {
	return func(o *Options) {
		o.maxlevel = l
	}
}
