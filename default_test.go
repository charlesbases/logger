package logger

import (
	"os"
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	SetDefault(WithWriter(os.Stdout))
	Debug(time.Now().Unix())
}
