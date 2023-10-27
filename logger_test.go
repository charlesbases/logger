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

func TestDefault(t *testing.T) {
	Debug("none")

	SetDefault(func(o *Options) {
		o.Name = "default"
	})
	Debug("default")

	named := Named("writer")
	named.Debug("default with file writer")
}

func TestCaller(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		Info("base")
		baseCallerSkip()

		a := Named("a")
		a.Info("a")

		b := a.Named("b")
		loggerCallerSkip(b)
	})

	t.Run("default", func(t *testing.T) {
		SetDefault(func(o *Options) {
			o.Name = "default"
			o.ContextHook = hook
		})

		Info("default")
		baseCallerSkip()
		WithContext(ctx).Info("context hook")

		a := Named("a")
		a.Info("a")
		a.WithContext(ctx).Info("context hook")

		b := a.Named("b")
		loggerCallerSkip(b)
		b.WithContext(ctx).Info("context hook")
	})
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

	b.Run("base", func(b *testing.B) {
		b.Run("named", func(b *testing.B) {
			bench(func() {
				Named("a").Info("a")
			})
		})
		b.Run("caller", func(b *testing.B) {
			bench(func() {
				baseCallerSkip()
			})
		})
	})
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

	SetDefault(func(o *Options) {
		o.Name = "default"
		o.ContextHook = hook
	})
	b.Run("default", func(b *testing.B) {
		b.Run("named", func(b *testing.B) {
			bench(func() {
				Named("a").Info("a")
			})
		})
		b.Run("caller", func(b *testing.B) {
			bench(func() {
				baseCallerSkip()
			})
		})
		b.Run("context", func(b *testing.B) {
			bench(func() {
				WithContext(ctx).Info("default")
			})
		})
	})
}

// Benchmark    	    1704	    875271 ns/op	  126242 B/op	    2506 allocs/op
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

	SetDefault(func(o *Options) {
		o.Name = "default"
		o.ContextHook = hook
	})

	bench(func() {
		Named("a").Info("a")
		WithContext(ctx).Info("ctx")
	})
}
