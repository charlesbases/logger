package zap

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var loop int = 1e4

	Init(WithService("zap")) // 830.113708ms

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
