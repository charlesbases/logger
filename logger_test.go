package logger

import (
	"fmt"
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

func TestNamed(t *testing.T) {
	// default
	{
		// SetDefault(func(o *Options) {
		// 	o.Name = "A"
		// 	o.Writer = filewriter.New(filewriter.OutputPath("log.log"))
		// })
		// Debugf("A: %p", base)
		// Flush()
		//
		// b := Named("B")
		// b.Debugf("B: %p", b)
		// b.Flush()
		//
		// c := Named("C")
		// c.Debugf("C: %p", c)
		// c.Flush()
		//
		// d := c.Named("D")
		// d.Debugf("D: %p", d)
		// d.Flush()
	}

	// new
	{
		// a := New(func(o *Options) {
		// 	o.Name = "AA"
		// 	o.Writer = filewriter.New(filewriter.OutputPath("log.log"))
		// })
		// a.Debugf("AA: %p", base)
		// a.Flush()
		//
		// b := a.Named("BB")
		// b.Debugf("BB: %p", b)
		// b.Flush()
		//
		// c := b.Named("CC")
		// c.Debugf("CC: %p", c)
		// c.Flush()
	}

	// warp
	{
		print(1) // print current line
	}

	// {
	// 	var loop int = 1e4
	//
	// 	var start = time.Now()
	// 	for i := 0; i < loop; i++ {
	// 		b := Named("B")
	// 		b.Debugf("B: %p", b)
	// 	}
	// 	fmt.Println(time.Since(start))
	// }
}

// print .
func print(skip int) {
	a := Named("A", func(o *Options) {
		o.Skip = skip
	})
	a.Debug("A")
	a.Flush()

	b := a.Named("B", func(o *Options) {
		o.Skip = skip
	})
	b.Error("B")
	b.Flush()
}

func TestFileWrite(t *testing.T) {
	// 并发写入同一个日志文件
	a := New(func(o *Options) {
		o.Name = "A"
		o.Writer = filewriter.New(filewriter.OutputPath("log.log"))

	})
	b := New(func(o *Options) {
		o.Name = "B"
		o.Writer = filewriter.New(filewriter.OutputPath("log.log"))
	})
	c := New(func(o *Options) {
		o.Name = "C"
		o.Writer = filewriter.New(filewriter.OutputPath("log.log"))
	})

	var count = 1 << 10
	var swg sync.WaitGroup
	swg.Add(3)

	go func() {
		for i := 0; i < count; i++ {
			a.Info("aaaaaaaaaa")
		}
		a.Flush()
		swg.Done()
	}()

	go func() {
		for i := 0; i < count; i++ {
			b.Debug("bbbbbbbbbb")
		}
		b.Flush()
		swg.Done()
	}()

	go func() {
		for i := 0; i < count; i++ {
			c.Error("cccccccccc")
		}
		c.Flush()
		swg.Done()
	}()

	swg.Wait()

	time.Sleep(time.Second * 3)
}

// now .
func now() string {
	return time.Now().Format(defaultDateFormat)
}
