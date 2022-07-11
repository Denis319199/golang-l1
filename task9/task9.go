package task9

import (
	"context"
	"fmt"
	"sync"
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

func Plus2AndPrint(in <-chan int, ctx context.Context, workersCount int) <-chan struct{} {
	wait := sync.WaitGroup{}
	wait.Add(workersCount)

	// Many workers for this task is redundant, since we use Println()
	for i := 0; i < workersCount; i++ {
		go func() {
			defer wait.Done()

			for {
				select {
				case val, ok := <-in:
					if ok {
						fmt.Println(val + 2)
					} else {
						return
					}
				case <-ctx.Done():
					return
				}
			}

		}()
	}

	done := make(chan struct{})

	go func() {
		wait.Wait()
		close(done)
	}()

	return done
}

func Task9() {
	// The same implementation as in task 2
	seq := GenerateSequence(1, 100, 1, context.Background())
	<-Plus2AndPrint(seq, context.Background(), 1)
}
