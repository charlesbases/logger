package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/charlesbases/logger/filewriter"
)

func TestMultipleLogger(t *testing.T) {
	Debug("Base logger")

	SetDefault(Service("SetDefault"), Writer(filewriter.New(filewriter.FilePath("./log/default.log"))))
	Debug("SetDefault logger")

	newl := New(Service("NEW"), Writer(filewriter.New(filewriter.FilePath("./log/default.log"))))
	newl.Debug("New logger")

	newl.Flush()
}

func TestNewLogger(t *testing.T) {
	var loop int = 1e4

	logger := New(Service("NEW"), Writer())

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

func TestSetDefault(t *testing.T) {
	var loop int = 1e4

	SetDefault(Service("SetDefault"), Writer())

	var start = time.Now()
	for i := 0; i < loop; i++ {
		Debug(now())
		Info(now())
		Warn(now())
		Error(now())
	}
	fmt.Println(time.Since(start))

	Flush()
}

func Test(t *testing.T) {
	a := New(Service("A"))
	a.Debugf("A: %p", base)

	{
		b := Name("B")
		b.Debugf("B: %p", b)
	}

	{
		var loop int = 1e4

		var start = time.Now()
		for i := 0; i < loop; i++ {
			b := Name("B")
			b.Debugf("B: %p", b)
		}
		fmt.Println(time.Since(start))
	}

	<-time.NewTicker(time.Second).C
}

// now .
func now() string {
	return time.Now().Format(DefaultDateFormat)
}
