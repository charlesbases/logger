package logger

import (
	"github.com/charlesbases/colors"
	"go.uber.org/zap/zapcore"
)

var (
	debugEncoder   = newEncoder("DBG", colors.PurpleSprint)
	infoEncoder    = newEncoder("INF", colors.GreenSprint)
	warnEncoder    = newEncoder("WRN", colors.BlueSprint)
	errorEncoder   = newEncoder("ERR", colors.RedSprint)
	fatalEncoder   = newEncoder("FAT", colors.RedSprint)
	unknownEncoder = newEncoder("UNK", colors.WhiteSprint)
)

// levelEncoder .
type levelEncoder struct {
	origin   string
	colorful string
}

// newEncoder .
func newEncoder(origin string, encoder func(a ...interface{}) string) *levelEncoder {
	return &levelEncoder{origin: origin, colorful: encoder(origin)}
}

// getEncoder .
func getEncoder(code zapcore.Level) *levelEncoder {
	switch code {
	case zapcore.DebugLevel:
		return debugEncoder
	case zapcore.InfoLevel:
		return infoEncoder
	case zapcore.WarnLevel:
		return warnEncoder
	case zapcore.ErrorLevel:
		return errorEncoder
	case zapcore.FatalLevel:
		return fatalEncoder
	default:
		return unknownEncoder
	}
}

// originEncoder .
func originEncoder(code zapcore.Level) string {
	return getEncoder(code).origin
}

// colorfulEncoder .
func colorfulEncoder(code zapcore.Level) string {
	return getEncoder(code).colorful
}
