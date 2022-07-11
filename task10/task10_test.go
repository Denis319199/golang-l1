package task10

import (
	"math/rand"
	"os"
	"testing"
)

var arr [100000000]float32

func TestMain(m *testing.M) {
	for i := 0; i < 100000000; i++ {
		arr[i] = rand.Float32()
	}

	os.Exit(m.Run())
}

// BenchmarkGroupBy-8                     1        1803871400 ns/op
func BenchmarkGroupBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GroupBy(arr[:], GetTemperatureGroup)
	}
}

// BenchmarkGroupByConcurrent-8           2         854193850 ns/op
func BenchmarkGroupByConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GroupByConcurrent(arr[:], GetTemperatureGroup)
	}
}
