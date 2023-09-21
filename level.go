package logger

import (
	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

type level int8

const (
	debugLevel level = iota - 1
	infoLevel
	warnLevel
	errorLevel
	fatalLevel

	minlevel = debugLevel
)

var render = map[zapcore.Level]func(b bool) string{
	zapcore.DebugLevel: func(b bool) string {
		if b {
			return colors.PurpleSprint("DBG")
		}
		return "DBG"
	},
	zapcore.InfoLevel: func(b bool) string {
		if b {
			return colors.GreenSprint("INF")
		}
		return "INF"
	},
	zapcore.WarnLevel: func(b bool) string {
		if b {
			return colors.BlueSprint("WRN")
		}
		return "WRN"
	},
	zapcore.ErrorLevel: func(b bool) string {
		if b {
			return colors.RedSprint("ERR")
		}
		return "ERR"
	},
	zapcore.FatalLevel: func(b bool) string {
		if b {
			return colors.RedSprint("FAT")
		}
		return "FAT"
	},
}

var string2level = map[string]level{
	"debug": debugLevel,
	"info":  infoLevel,
	"warn":  warnLevel,
	"error": errorLevel,
	"fatal": fatalLevel,
}

// shortName .
func shortName(zl zapcore.Level) func(b bool) string {
	if fn, found := render[zl]; found {
		return fn
	}
	return func(b bool) string {
		if b {
			return colors.RedSprint("UNK")
		}
		return "UNK"
	}
}
