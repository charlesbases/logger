package logger

import (
	"fmt"

	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

type level int8

const (
	TraceLevel level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel

	minlevel = TraceLevel
	maxlevel = FatalLevel
)

var render = map[level]string{
	TraceLevel: colors.WhiteSprint("TRC"),
	DebugLevel: colors.PurpleSprint("DBG"),
	InfoLevel:  colors.GreenSprint("INF"),
	WarnLevel:  colors.BlueSprint("WRN"),
	ErrorLevel: colors.RedSprint("ERR"),
	FatalLevel: colors.RedSprint("FAT"),
}

var zap2level = map[zapcore.Level]level{
	zapcore.DebugLevel:  DebugLevel,
	zapcore.InfoLevel:   InfoLevel,
	zapcore.WarnLevel:   WarnLevel,
	zapcore.ErrorLevel:  ErrorLevel,
	zapcore.DPanicLevel: FatalLevel,
	zapcore.PanicLevel:  FatalLevel,
	zapcore.FatalLevel:  FatalLevel,
}

var string2level = map[string]level{
	"trace": TraceLevel,
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"fatal": FatalLevel,
}

// convertZapLevel convert zapcore.Level to level
func convertZapLevel(lv zapcore.Level) level {
	if l, found := zap2level[lv]; found {
		return l
	}
	return TraceLevel
}

// convertString .
func convertString(v string) level {
	if l, found := string2level[v]; found {
		return l
	}
	fmt.Println(fmt.Sprintf(`unknown logger level(%s). use default.`, v))
	return TraceLevel
}
