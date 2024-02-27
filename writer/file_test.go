package writer

import (
	"os"
	"testing"
)

func BenchmarkStderr(b *testing.B) {
	var bench = func(f func()) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			f()
		}
		b.StopTimer()
	}

	bench(
		func() {
			stderr(os.ErrNotExist)
		},
	)
}
