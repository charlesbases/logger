package logger

import (
	"os"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	w := NewFileWriter(
		WithFileDir("./logs"),
		WithFileName("app.log"),
		WithMaxAge(7),
		WithMaxSize(512),
		WithCompress(true),
	)

	l := New(WithWriter(os.Stdout, w))
	l.Debug("now:", time.Now().UTC().Format(time.DateTime))
}

//go:generate go test -run BenchmarkNew -test.bench=BenchmarkNew -test.benchmem .
func BenchmarkNew(b *testing.B) {
	var bench = func(f func()) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			f()
		}
		b.StopTimer()
	}

	l := New()

	bench(func() {
		l.Debug("now:", time.Now().UTC().Format(time.DateTime))
		l.Info("now:", time.Now().UTC().Format(time.DateTime))
		l.Warn("now:", time.Now().UTC().Format(time.DateTime))
		l.Error("now:", time.Now().UTC().Format(time.DateTime))
	})
}
