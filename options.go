package logger

import "strings"

const (
	// DefaultSkip .
	DefaultSkip = 1
	// DefaultMaxRolls 日志保留时间
	DefaultMaxRolls = 7
	// DefaultDateFormat date format
	DefaultDateFormat = "2006-01-02 15:04:05.000"
	// DefaultFilePath default file path
	DefaultFilePath = "./log/log.log"
)

// Options .
type Options struct {
	// Service 服务名
	Service string
	// FilePath 日志文件路径
	FilePath string
	// MaxRolls 日志保留天数
	MaxRolls int
	// MinLevel 允许的最小日志级别. default: "Trace"
	MinLevel string
	// MaxLevel 允许的最大日志级别. default: "Fatal"
	MaxLevel string
	// Skip 跳过的调用者数量. default: DefaultSkip
	Skip int

	// minlevel convert from MinLevel. default: _minlevel
	minlevel level
	// maxlevel convert from MaxLevel. default: _maxlevel
	maxlevel level
}

// defaultOption .
func defaultOption() *Options {
	return &Options{
		FilePath: DefaultFilePath,
		MaxRolls: DefaultMaxRolls,
		Skip:     DefaultSkip,

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

// WithFilePath .
func WithFilePath(filename string) Option {
	return func(o *Options) {
		o.FilePath = filename
	}
}

// WithMaxRolls .
func WithMaxRolls(rolls int) Option {
	return func(o *Options) {
		o.MaxRolls = rolls
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
