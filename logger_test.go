package logger

import (
	"context"
	"sync"
	"testing"
	"time"
)

// now .
func now() string {
	return time.Now().Format(defaultDateFormat)
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
	// default
	Info("default")

	// SetDefault
	SetDefault()
	Info("SetDefault")

	// Named
	a := Named("a")
	a.Info("a")

	b := a.Named("b")
	b.Info("b")

	// CakkerSkip
	testCallerSkip(b)
}

// testCallerSkip .
func testCallerSkip(l *Logger) {
	l.Named("c").CallerSkip(1).Info("c")
}

func TestContextHook(t *testing.T) {
	SetDefault(func(o *Options) {
		o.ContextHook = func(ctx context.Context) func(l *Logger) *Logger {
			return func(l *Logger) *Logger {
				if traceid, ok := ctx.Value("traceid").(string); ok && len(traceid) != 0 {
					return l.Named(traceid)
				}
				return l
			}
		}
	})

	ctx := context.WithValue(context.Background(), "traceid", "123456")
	WithContext(ctx).Info(time.Now())
}

// BenchmarkLogger-16    	  110325	     12282 ns/op	    1824 B/op	      36 allocs/op
func BenchmarkLogger(b *testing.B) {
	var bench = func(f func()) {
		b.ResetTimer()
		f()
		b.StopTimer()
	}

	bench(func() {
		wg := sync.WaitGroup{}
		wg.Add(3)

		go func() {
			for i := 0; i < b.N; i++ {
				Named("a").Info("aa")
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < b.N; i++ {
				Named("b").Info("bb")
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < b.N; i++ {
				Named("c").Info("cc")
			}
			wg.Done()
		}()

		wg.Wait()
	})
}
