package logger

import "io"

const (
	// defaultCallerSkip .
	defaultCallerSkip = 1
	// defaultDateFormat date format
	defaultDateFormat = "2006-01-02 15:04:05.000"
)

// Options .
type Options struct {
	// Name 名称
	Name string
	// Skip 跳过的调用者数量. default: defaultCallerSkip
	Skip int
	// MinLevel 允许的最小日志级别. default: "trace"
	MinLevel string
	// MaxLevel 允许的最大日志级别. default: "fatal"
	MaxLevel string
	// Writer others output
	Writer io.Writer

	baseSkip int

	// minlevel convert MinLevel to level
	minlevel level
	// maxlevel convert from MaxLevel to level
	maxlevel level
}

// defaultOptions .
func defaultOptions() *Options {
	return &Options{
		Skip: defaultCallerSkip,

		minlevel: minlevel,
		maxlevel: maxlevel,
	}
}
