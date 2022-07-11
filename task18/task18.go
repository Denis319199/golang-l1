package task18

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	Value int32
}

func (counter *AtomicCounter) Increment() int32 {
	return atomic.AddInt32(&counter.Value, 1)
}

func (counter *AtomicCounter) Decrement() int32 {
	return atomic.AddInt32(&counter.Value, -1)
}

func Task18() {
	counter := AtomicCounter{}

	wait := sync.WaitGroup{}
	wait.Add(30)
	for i := 0; i < 30; i++ {

		go func() {
			counter.Increment()
			wait.Done()
		}()
	}

	wait.Wait()
	// I suppose Done() and Wait() calls create a synchronizes-with relationship
	fmt.Println(counter.Value) // SHOULD I DO STORE?
}
