package task2

import (
	"context"
	"fmt"
	"sync"
)

// GenerateSequence - just puts values into channel returned
func GenerateSequence(start, end, step int, ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for val := start; val < end; val += step {
			select {
			case out <- val:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// Pow - creates 'workersCount' goroutines which take values from the given channel,
// calculate square of each one and put the result into channel returned
func Pow(in <-chan int, ctx context.Context, workersCount int) <-chan int {
	out := make(chan int)

	wait := sync.WaitGroup{}
	wait.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			defer wait.Done()

			for val := range in {
				select {
				case out <- val * val:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wait.Wait()
		close(out)
	}()

	return out
}

func Task2() {
	// Implemented pipeline with cancellation
	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	for val := range Pow(GenerateSequence(2, 100, 2, ctx), ctx, 2) {
		fmt.Println(val)
	}
}
