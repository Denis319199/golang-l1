package task4

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Run(workersCount int, ctx context.Context) <-chan struct{} {
	doneChan := make(chan struct{})

	go func() {
		c := make(chan string)

		wait := sync.WaitGroup{}
		wait.Add(workersCount)

		for i := 0; i < workersCount; i++ {
			go func() {
				// Reads values until the channel has been closed
				for val := range c {
					fmt.Println(val)
					time.Sleep(time.Millisecond)
				}

				wait.Done()
			}()
		}

		for {
			select {
			case c <- RandStringRunes(20):
			case <-ctx.Done():
				// The order is important
				close(c)
				wait.Wait() // SHOULD I AWAIT??? - I think YES

				// I can just close channel and get the same outcome
				doneChan <- struct{}{} // That is, I can remove this line
				close(doneChan)
			}
		}
	}()

	return doneChan
}

func Task4() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// Runs 10 read-workers + 1 auxiliary write-worker
	doneChan := Run(10, ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		fmt.Println("SIGNAL")
	case <-ctx.Done():
		fmt.Println("TIMEOUT")
	}

	cancel()

	// SHOULD I DO THIS??? - YES
	// awaits all workers to be terminated
	<-doneChan
}
