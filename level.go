package logger

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

type level int8

const (
	levelTrace level = iota
	levelDebug
	levelInfo
	levelWarn
	levelError
	levelFatal

	_minlevel = levelTrace
	_maxlevel = levelFatal
)

type attribute int

const (
	reset  = 0
	escape = "\x1b"
)

const (
	black attribute = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

var colors = [6]attribute{}

var shorts = [6]string{}

var convert = map[zapcore.Level]level{
	zapcore.DebugLevel:  levelDebug,
	zapcore.InfoLevel:   levelInfo,
	zapcore.WarnLevel:   levelWarn,
	zapcore.ErrorLevel:  levelError,
	zapcore.DPanicLevel: levelFatal,
	zapcore.PanicLevel:  levelFatal,
	zapcore.FatalLevel:  levelFatal,
}

func init() {
	colors[levelTrace] = white
	colors[levelDebug] = magenta
	colors[levelInfo] = green
	colors[levelWarn] = blue
	colors[levelError] = red
	colors[levelFatal] = red

	shorts[levelTrace] = levelTrace.render()
	shorts[levelDebug] = levelDebug.render()
	shorts[levelInfo] = levelInfo.render()
	shorts[levelWarn] = levelWarn.render()
	shorts[levelError] = levelError.render()
	shorts[levelFatal] = levelFatal.render()
}

// render .
func (l level) render() string {
	return fmt.Sprintf("%s[%dm%s%s[%dm", escape, l.color(), l.string(), escape, reset)
}

// color .
func (l level) color() attribute {
	return colors[l]
}

// short .
func (l level) short() string {
	return shorts[l]
}

// string .
func (l level) string() string {
	switch l {
	case levelTrace:
		return "TRC"
	case levelDebug:
		return "DBG"
	case levelInfo:
		return "INF"
	case levelWarn:
		return "WRN"
	case levelError:
		return "ERR"
	case levelFatal:
		return "FAT"
	default:
		return fmt.Sprintf("UNK[%d]", l)
	}
}

// convertZapLevel zapcore.Level to level
func convertZapLevel(lv zapcore.Level) level {
	if l, existing := convert[lv]; existing {
		return l
	}
	return levelTrace
}
