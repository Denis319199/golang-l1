package task3

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

// Sum - in fact, the implementation is the same as for Pow
func Sum(in <-chan int, ctx context.Context, workersCount int) <-chan int {
	out := make(chan int)

	go func() {
		workerResults := make(chan int)

		wait := sync.WaitGroup{}
		wait.Add(workersCount)

		for i := 0; i < workersCount; i++ {
			go func() {
				defer wait.Done()

				// I can wait for the 'in' channel to be closed, but I don't do
				// so because I want goroutines to stop instantly when context is done

				// Added some delay for benchmarks
				sum := 0
				for {
					select {
					case val, ok := <-in:
						if ok {
							sum += val
							time.Sleep(time.Millisecond) // Required for benchmarks
						} else {
							workerResults <- sum
							return
						}

					case <-ctx.Done():
						workerResults <- sum
						return
					}
				}
			}()
		}

		go func() {
			wait.Wait()
			close(workerResults)
		}()

		sum := 0

		// Too verbose but can instantly end tasks processing (I think
		// it is appropriate for long tasks)
		/*for val := range workerResults {
			sum += val
			select {
			case val, ok := <-workerResults:
				if ok {

				} else {
					out <- sum
					return
				}
			case <-ctx.Done():
				out <- sum
				return
			}
		}*/

		for val := range workerResults {
			sum += val
		}

		out <- sum
		close(out)
	}()

	return out
}

func Task3() {
	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	fmt.Println(<-Sum(Pow(GenerateSequence(2, 1000, 1, ctx), ctx, 2), ctx, 2))
}
