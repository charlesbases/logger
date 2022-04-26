package logger

import (
	"strings"

	"github.com/charlesbases/logger/filewriter"
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
	// FileWriterOption 文件写入配置
	FileWriterOptions []filewriter.Option
	// MinLevel 允许的最小日志级别. default: "Trace"
	MinLevel string
	// MaxLevel 允许的最大日志级别. default: "Fatal"
	MaxLevel string
	// Skip 跳过的调用者数量. default: DefaultSkip
	Skip int

	// 是否写入文件
	store bool
	// minlevel convert from MinLevel. default: _minlevel
	minlevel level
	// maxlevel convert from MaxLevel. default: _maxlevel
	maxlevel level
}

// defaultOption .
func defaultOption() *Options {
	return &Options{
		Skip: DefaultSkip,

		store:    false,
		minlevel: _minlevel,
		maxlevel: _maxlevel,
	}
}

type Option func(o *Options)

// WithSkip .
func WithSkip(skip int) Option {
	return func(o *Options) {
		o.Skip = skip
	}
}

// WithService .
func WithService(service string) Option {
	return func(o *Options) {
		o.Service = service
	}
}

// WithFileWriter .
func WithFileWriter(opts ...filewriter.Option) Option {
	return func(o *Options) {
		o.store = true
		o.FileWriterOptions = opts
	}
}

// WithMinLevel allowed: trace | debug | info | warn | error | fatal
func WithMinLevel(l string) Option {
	return func(o *Options) {
		switch strings.ToLower(l) {
		case "trace":
			o.minlevel = levelTrace
		case "debug":
			o.minlevel = levelDebug
		case "info":
			o.minlevel = levelInfo
		case "warn":
			o.minlevel = levelWarn
		case "error":
			o.minlevel = levelError
		case "fatal":
			o.minlevel = levelFatal
		}
	}
}

// WithMaxLevel allowed: trace | debug | info | warn | error | fatal
func WithMaxLevel(l string) Option {
	return func(o *Options) {
		switch strings.ToLower(l) {
		case "trace":
			o.maxlevel = levelTrace
		case "debug":
			o.maxlevel = levelDebug
		case "info":
			o.maxlevel = levelInfo
		case "warn":
			o.maxlevel = levelWarn
		case "error":
			o.maxlevel = levelError
		case "fatal":
			o.maxlevel = levelFatal
		}
	}
}
