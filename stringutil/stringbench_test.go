package stringutil

import (
	"testing"
)

func Benchmark_GetNewId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetNewId()
	}
}
