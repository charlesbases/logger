package logger

import (
	"fmt"
	"testing"
	"time"
)

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
