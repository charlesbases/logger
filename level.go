package logger

import (
	"strings"

	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

const (
	traceLevel level = iota
	debugLevel
	infoLevel
	warnLevel
	errorLevel
	fatalLevel

	minlevel = traceLevel
	maxlevel = fatalLevel
)

type level int8

var render = map[level]string{
	traceLevel: colors.WhiteSprint("TRC"),
	debugLevel: colors.PurpleSprint("DBG"),
	infoLevel:  colors.GreenSprint("INF"),
	warnLevel:  colors.BlueSprint("WRN"),
	errorLevel: colors.RedSprint("ERR"),
	fatalLevel: colors.RedSprint("FAT"),
}

var convert = map[zapcore.Level]level{
	zapcore.DebugLevel:  debugLevel,
	zapcore.InfoLevel:   infoLevel,
	zapcore.WarnLevel:   warnLevel,
	zapcore.ErrorLevel:  errorLevel,
	zapcore.DPanicLevel: fatalLevel,
	zapcore.PanicLevel:  fatalLevel,
	zapcore.FatalLevel:  fatalLevel,
}

// convertZapLevel zapcore.Level to level
func convertZapLevel(lv zapcore.Level) level {
	if l, existing := convert[lv]; existing {
		return l
	}
	return traceLevel
}

// convertString .
func convertString(s string) level {
	switch strings.ToLower(s) {
	case "trace":
		return traceLevel
	case "debug":
		return debugLevel
	case "info":
		return infoLevel
	case "warn":
		return warnLevel
	case "error":
		return errorLevel
	case "fatal":
		return fatalLevel
	default:
		return -1
	}
}
