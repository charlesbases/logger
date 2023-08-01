package logger

import (
	"fmt"

	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

type level int8

const (
	trace level = iota
	debug
	info
	warn
	error
	fatal

	minlevel = trace
	maxlevel = fatal
)

var render = map[level]string{
	trace: colors.WhiteSprint("TRC"),
	debug: colors.PurpleSprint("DBG"),
	info:  colors.GreenSprint("INF"),
	warn:  colors.BlueSprint("WRN"),
	error: colors.RedSprint("ERR"),
	fatal: colors.RedSprint("FAT"),
}

var zap2level = map[zapcore.Level]level{
	zapcore.DebugLevel:  debug,
	zapcore.InfoLevel:   info,
	zapcore.WarnLevel:   warn,
	zapcore.ErrorLevel:  error,
	zapcore.DPanicLevel: fatal,
	zapcore.PanicLevel:  fatal,
	zapcore.FatalLevel:  fatal,
}

var string2level = map[string]level{
	"trace": trace,
	"debug": debug,
	"info":  info,
	"warn":  warn,
	"error": error,
	"fatal": fatal,
}

// convertZapLevel convert zapcore.Level to level
func convertZapLevel(lv zapcore.Level) level {
	if l, found := zap2level[lv]; found {
		return l
	}
	return trace
}

// convertString .
func convertString(v string) level {
	if l, found := string2level[v]; found {
		return l
	}
	fmt.Println(fmt.Sprintf(`unknown logger level(%s). use default.`, v))
	return trace
}
