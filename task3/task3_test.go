package task3

import (
	"fmt"
	"testing"
	"time"
)

// Conclusion: concurrent processing (2 workers) tasks taking at least 1ms is
// 1.999 times faster than non-concurrent

// BenchmarkTask3-8               1        7735705800 ns/op
func BenchmarkTask3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Task3()
	}
}

// BenchmarkTask3Plain-8                  1        15464522800 ns/op
func BenchmarkTask3Plain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 2; i < 1000; i++ {
				out <- i
			}
		}()
		sum := 0
		for val := range out {
			sum += val * val
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sum)
	}
}

// A bit faster than the function above due to not using channel to send values,
// but it is approximately as slow as the function above relative to the first option
// BenchmarkTask3PlainPlain-8             1        15293118800 ns/op
func BenchmarkTask3PlainPlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for i := 2; i < 1000; i++ {
			sum += i * i
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sum)
	}
}
