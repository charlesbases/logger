package logger

import (
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	l := New(WithWriter(os.Stdout))
	l.Debug("now:", time.Now().UTC().Format(time.DateTime))
}
