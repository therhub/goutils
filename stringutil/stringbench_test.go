package stringutil

import (
	"testing"
)

func Benchmark_TryCreateNewId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TryCreateNewId()
	}
}
