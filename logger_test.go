package logger

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/charlesbases/logger/filewriter"
)

// now .
func now() string {
	return time.Now().Format(defaultDateFormat)
}

func TestMultipleLogger(t *testing.T) {
	Debug("Base logger")

	SetDefault(func(o *Options) {
		o.Name = "Default"
		o.Writer = filewriter.New()
	})
	Debug("SetDefault logger")

	newl := New(func(o *Options) {
		o.Name = "New"
		o.Writer = filewriter.New()
	})
	newl.Debug("New logger")

	newl.Flush()
}

func TestNewLogger(t *testing.T) {
	var loop int = 1e4

	logger := New(func(o *Options) {
		o.Name = "New"
	})

	var start = time.Now()
	for i := 0; i < loop; i++ {
		logger.Debug(now())
		logger.Info(now())
		logger.Warn(now())
		logger.Error(now())
	}
	fmt.Println(time.Since(start))

	<-time.After(time.Second * 1)
}

func TestCaller(t *testing.T) {
	// default
	{
		SetDefault(func(o *Options) { o.Name = "Default" })
		Debug(59)
		CallerSkip(0).Info(60)

		a := Named("A")
		a.Debug(63)
		a.CallerSkip(0).Info(64)

		b := a.Named("B")
		b.Debug(67)
		b.CallerSkip(0).Info(68)
	}

	// new
	{
		n := New(func(o *Options) { o.Name = "New" })
		n.Error(74)
		n.CallerSkip(0).Info(75)

		a := n.Named("A")
		a.Error(78)
		a.CallerSkip(0).Info(79)

		b := a.Named("B")
		b.Error(82)
		b.CallerSkip(0).Info(83)
	}

	print(86)
}

// print .
func print(line int) {
	a := Named("C", func(o *Options) { o.Skip = 1 })
	a.Warn(line)

	b := a.Named("", func(o *Options) { o.Skip = 1 })
	b.Warn(line)
}

func TestFileWriter(t *testing.T) {
	SetDefault(func(o *Options) {
		o.Writer = filewriter.New()
	})

	for i := 0; i < 10; i++ {
		go func() {
			tk := time.NewTicker(time.Second)
			for {
				select {
				case <-tk.C:
					Info(now())
				}
			}
		}()
	}

	select {}
}

func TestBase(t *testing.T) {
	var loop int = 1e4

	var start = time.Now()
	for i := 0; i < loop; i++ {
		Debug(now())
		Info(now())
		Warn(now())
		Error(now())
	}
	fmt.Println(time.Since(start))

	<-time.After(time.Second * 1)
}

func TestContextHook(t *testing.T) {
	log := New(func(o *Options) {
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
	log.WithContext(ctx).Info(time.Now())
}

func TestTime(t *testing.T) {
	// samples: 10
	// minimum: 291.569ms
	// maximum: 424.6169ms
	// average: 309.05249ms
	bench(func(i int) {
		log := Named(strconv.Itoa(i))
		log.Info(i)
	})
}

func bench(fn func(id int)) {
	var number, count = 10000, 10

	var min, max, total time.Duration
	for i := 0; i < count; i++ {
		start := time.Now()
		for i := 0; i < number; i++ {
			fn(i)
		}
		sub := time.Since(start)

		switch {
		case max == 0:
			max = sub
			min = sub
		case sub > max:
			max = sub
		case sub < min:
			min = sub
		}

		total += sub
	}

	fmt.Println("samples:", count)
	fmt.Println("minimum:", min)
	fmt.Println("maximum:", max)
	fmt.Println("average:", total/time.Duration(count))
	fmt.Println()
}
