package logger

import (
	"fmt"
	"testing"
	"time"

	"github.com/charlesbases/logger/filewriter"
)

func TestMultipleLogger(t *testing.T) {
	Debug("Base logger")

	SetDefault(WithService("SetDefault"), WithFileWriter(filewriter.FilePath("./log/default.log")))
	Debug("SetDefault logger")

	newl := New(WithService("NEW"), WithFileWriter(filewriter.FilePath("./log/new.log")))
	newl.Debug("New logger")

	<-time.After(time.Second * 1)
}

func TestNewLogger(t *testing.T) {
	var loop int = 1e4

	logger := New(WithService("NEW"), WithFileWriter())

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

	SetDefault(WithService("SetDefault"), WithFileWriter())

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

// now .
func now() string {
	return time.Now().Format(DefaultDateFormat)
}
