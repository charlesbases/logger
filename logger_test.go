package logger

import (
	"context"
	"sync"
	"testing"
	"time"
)

var hook ContextHook = func(ctx context.Context) string {
	return ctx.Value("traceid").(string)
}

var ctx = context.WithValue(context.Background(), "traceid", "traceid")

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
	SetDefault(WithName("default"), WithContextHook(hook), WithColorful())
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
// (non-filewrite)    	    2824	    451700 ns/op	  102554 B/op	    2206 allocs/op
//
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

	SetDefault(WithName("default"), WithContextHook(hook), WithColorful())

	bench(
		func() {
			Named("a").Info("a")
			Context(ctx).Error("ctx")
		},
	)
}
