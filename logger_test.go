package logger

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/charlesbases/logger/filewriter"
)

func TestMultipleLogger(t *testing.T) {
	Debug("Base logger")

	SetDefault(func(o *Options) {
		o.Name = "Default"
		o.Writer = filewriter.New(filewriter.MaxRolls(7), filewriter.OutputPath("log.log"))
	})
	Debug("SetDefault logger")

	newl := New(func(o *Options) {
		o.Name = "New"
		o.Writer = filewriter.New(filewriter.MaxRolls(7), filewriter.OutputPath("log.log"))
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

func TestTime(t *testing.T) {
	var count, samples = 10000, 10

	fn := func(i int) {
		log := Named(strconv.Itoa(i))
		log.Info(i)
	}

	var min, max, total time.Duration
	for i := 0; i < samples; i++ {
		start := time.Now()
		for i := 0; i < count; i++ {
			fn(i)
		}
		sub := time.Since(start)
		fmt.Println(fmt.Sprintf("count %d: %v", i+1, sub))
		if sub > max {
			max = sub
		}
		if sub < min || min == 0 {
			min = sub
		}
		total += sub
	}

	fmt.Println("minimum:", min)      // 344.2704ms
	fmt.Println("maximum:", max)      // 461.7053ms
	fmt.Println("average:", total/10) // 382.72425ms
}

func TestCaller(t *testing.T) {
	// default
	{
		SetDefault(func(o *Options) { o.Name = "Default" })
		Debug(now())
		CallerSkip(0).Info(82)

		a := Named("A")
		a.Debug(85)
		a.CallerSkip(0).Info(86)

		b := a.Named("B")
		b.Debug(89)
		b.CallerSkip(0).Info(90)
	}

	// new
	{
		n := New(func(o *Options) { o.Name = "New" })
		n.Error(now())
		CallerSkip(0).Info(97)

		a := n.Named("A")
		a.Error(100)
		a.CallerSkip(0).Info(101)

		b := a.Named("B")
		b.Error(104)
		b.CallerSkip(0).Info(105)

	}

	print(109)
}

// print .
func print(line int) {
	a := Named("A", func(o *Options) { o.Skip = 1 })
	a.Warn(line)

	b := a.Named("B", func(o *Options) { o.Skip = 1 })
	b.Warn(line)
}

func TestFileWrite(t *testing.T) {
	SetDefault(func(o *Options) {
		o.Writer = filewriter.New(filewriter.OutputPath("log.log"))
	})

	var count = 10000
	var concurrency = 10

	var swg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		swg.Add(1)

		go func(ccy int) {
			name := string([]byte{byte(65 + ccy)})

			log := Named(name)
			var v string
			for i := 0; i < 10; i++ {
				v = v + name
			}

			fmt.Println(v)

			for i := 0; i < count; i++ {
				log.Info(v)
			}

			log.Flush()
			swg.Done()
		}(i)
	}

	swg.Wait()
}

// now .
func now() string {
	return time.Now().Format(defaultDateFormat)
}
