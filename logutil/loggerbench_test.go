package logutil

import "testing"

func Benchmark_write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SystemLog("info", "测试数据")
	}
}
