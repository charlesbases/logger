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

var renderFunc = func(s string, sf func(v ...interface{}) string) func(b bool) string {
	return func(b bool) string {
		if b {
			return sf(s)
		}
		return s
	}
}

var render = map[zapcore.Level]func(b bool) string{
	zapcore.DebugLevel: renderFunc("DBG", colors.PurpleSprint),
	zapcore.InfoLevel:  renderFunc("INF", colors.GreenSprint),
	zapcore.WarnLevel:  renderFunc("WRN", colors.BlueSprint),
	zapcore.ErrorLevel: renderFunc("ERR", colors.RedSprint),
	zapcore.FatalLevel: renderFunc("FAT", colors.RedSprint),
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
	return renderFunc("UNK", colors.RedSprint)
}
