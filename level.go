package logger

import (
	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

type Level int8

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	minlevel = TraceLevel
	maxlevel = FatalLevel
)

var render = map[Level]string{
	TraceLevel: colors.WhiteSprint("TRC"),
	DebugLevel: colors.PurpleSprint("DBG"),
	InfoLevel:  colors.GreenSprint("INF"),
	WarnLevel:  colors.BlueSprint("WRN"),
	ErrorLevel: colors.RedSprint("ERR"),
	FatalLevel: colors.RedSprint("FAT"),
}

var convert = map[zapcore.Level]Level{
	zapcore.DebugLevel:  DebugLevel,
	zapcore.InfoLevel:   InfoLevel,
	zapcore.WarnLevel:   WarnLevel,
	zapcore.ErrorLevel:  ErrorLevel,
	zapcore.DPanicLevel: FatalLevel,
	zapcore.PanicLevel:  FatalLevel,
	zapcore.FatalLevel:  FatalLevel,
}

// convertZapLevel zapcore.Level to Level
func convertZapLevel(lv zapcore.Level) Level {
	if l, existing := convert[lv]; existing {
		return l
	}
	return TraceLevel
}
