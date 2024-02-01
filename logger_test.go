package logger

import (
	"context"
	"sync"
	"testing"
	"time"
)

var hook ContextHook = func(ctx context.Context) func(l *Logger) *Logger {
	return func(l *Logger) *Logger {
		if ctx != context.Background() {
			return l.Named(ctx.Value("traceid").(string))
		}
		return l
	}
}

var ctx = context.WithValue(context.Background(), "traceid", "1")

// now .
func now() string {
	return time.Now().Format(defaultDateFormat)
}

// baseCallerSkip .
func baseCallerSkip() {
	CallerSkip(1).Info("caller skip")
}

// loggerCallerSkip .
func loggerCallerSkip(log *Logger) {
	log.CallerSkip(1).Info("caller skip")
}

func TestContextHook(t *testing.T) {
	SetDefault(WithName("default"), WithContextHook(hook))
	Context(ctx).Info("ctx")
}

func TestDefault(t *testing.T) {
	Debug("none")

	SetDefault(WithName("default"))
	Debug("default")

	named := Named("writer")
	named.Debug("default with file writer")
}

func TestCaller(t *testing.T) {
	t.Run(
		"defaultLogger", func(t *testing.T) {
			Info("defaultLogger")
			baseCallerSkip()

			a := Named("a")
			a.Info("a")

			b := a.Named("b")
			loggerCallerSkip(b)
		},
	)

	t.Run(
		"default", func(t *testing.T) {
			SetDefault(WithName("default"), WithContextHook(hook))

			Info("default")
			baseCallerSkip()
			Context(ctx).Info("context hook")

			a := Named("a")
			a.Info("a")
			a.Context(ctx).Info("context hook")

			b := a.Named("b")
			loggerCallerSkip(b)
			b.Context(ctx).Info("context hook")
		},
	)
}

func BenchmarkBase(b *testing.B) {
	var count = 100
	var bench = func(f func()) {
		b.ResetTimer()
		wg := sync.WaitGroup{}
		wg.Add(count)
		for idx := 0; idx < count; idx++ {
			go func() {
				for i := 0; i < b.N; i++ {
					f()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		b.StopTimer()
	}

	b.Run(
		"defaultLogger", func(b *testing.B) {
			b.Run(
				"named", func(b *testing.B) {
					bench(
						func() {
							Named("a").Info("a")
						},
					)
				},
			)
			b.Run(
				"caller", func(b *testing.B) {
					bench(
						func() {
							baseCallerSkip()
						},
					)
				},
			)
		},
	)
}

func BenchmarkDefault(b *testing.B) {
	var count = 100
	var bench = func(f func()) {
		b.ResetTimer()
		wg := sync.WaitGroup{}
		wg.Add(count)
		for idx := 0; idx < count; idx++ {
			go func() {
				for i := 0; i < b.N; i++ {
					f()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		b.StopTimer()
	}

	SetDefault(WithName("default"), WithContextHook(hook))
	b.Run(
		"default", func(b *testing.B) {
			b.Run(
				"named", func(b *testing.B) {
					bench(
						func() {
							Named("a").Info("a")
						},
					)
				},
			)
			b.Run(
				"caller", func(b *testing.B) {
					bench(
						func() {
							baseCallerSkip()
						},
					)
				},
			)
			b.Run(
				"context", func(b *testing.B) {
					bench(
						func() {
							Context(ctx).Info("default")
						},
					)
				},
			)
		},
	)
}

// (filewrite)       	     829	   1646762 ns/op	  129959 B/op	    3508 allocs/op
// (non-filewrite)    	    1686	    820628 ns/op	  103639 B/op	    2304 allocs/op
//
//go:generate go test -run Benchmark -test.bench=. -test.benchmem .
func Benchmark(b *testing.B) {
	var count = 100
	var bench = func(f func()) {
		b.ResetTimer()
		wg := sync.WaitGroup{}
		wg.Add(count)
		for idx := 0; idx < count; idx++ {
			go func() {
				for i := 0; i < b.N; i++ {
					f()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		b.StopTimer()
	}

	SetDefault(WithName("default"), WithContextHook(hook))

	bench(
		func() {
			Named("a").Info("a")
			Context(ctx).Info("ctx")
		},
	)
}
