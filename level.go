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

var renderFunc = func(sf func(v ...interface{}) string, v string) func(b bool) string {
	return func(b bool) string {
		if b {
			return sf(v)
		}
		return v
	}
}

var render = map[zapcore.Level]func(b bool) string{
	zapcore.DebugLevel: renderFunc(colors.PurpleSprint, "DBG"),
	zapcore.InfoLevel:  renderFunc(colors.GreenSprint, "INF"),
	zapcore.WarnLevel:  renderFunc(colors.BlueSprint, "WRN"),
	zapcore.ErrorLevel: renderFunc(colors.RedSprint, "ERR"),
	zapcore.FatalLevel: renderFunc(colors.RedSprint, "FAT"),
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
