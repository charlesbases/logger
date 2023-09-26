package filewriter

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestStderr(t *testing.T) {
	bench(func(id int) {
		stderr(os.ErrNotExist)
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
}
