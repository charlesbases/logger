package logger

import (
	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

type level int8

const (
	debug level = iota - 1
	info
	warn
	error
	fatal

	minlevel = debug
)

var render = map[zapcore.Level]string{
	zapcore.DebugLevel: colors.PurpleSprint("DBG"),
	zapcore.InfoLevel:  colors.GreenSprint("INF"),
	zapcore.WarnLevel:  colors.BlueSprint("WRN"),
	zapcore.ErrorLevel: colors.RedSprint("ERR"),
	zapcore.FatalLevel: colors.RedSprint("FAT"),
}

var string2level = map[string]level{
	"debug": debug,
	"info":  info,
	"warn":  warn,
	"error": error,
	"fatal": fatal,
}
