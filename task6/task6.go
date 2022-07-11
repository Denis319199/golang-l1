package task6

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// SomeRoutineContext - we can use chan struct{} or channel returned by time.After
// instead of Context but the essence is the same
func SomeRoutineContext(ctx context.Context) {
	defer fmt.Println("Done 1")

	// There is a minus: we should have a channel which will be filled with a result
	c := make(chan int)

	go func() {
		defer close(c)

		for i := 0; i < 10; i++ {
			select {
			case c <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	for {
		select {
		case val, ok := <-c:
			if ok {
				fmt.Println(val)
			} else {
				return
			}

		case <-ctx.Done():
			return
		}
	}
}

func SomeRoutineAtomic(atomicFlag *int32) {
	defer fmt.Println("Done 2")

	// Step 1
	time.Sleep(time.Second)

	// I hope GoLang supports release/acquired memory models or even sequential
	// consistency by default, so I can load() and store() without fear
	if atomic.LoadInt32(atomicFlag) == 1 {
		return
	}

	// Step 2
	time.Sleep(time.Second)
	if atomic.LoadInt32(atomicFlag) == 1 {
		return
	}

	// Step 3
	time.Sleep(time.Second)
}

// SomeRoutineChannel SomeRoutineTakingChannel - if values for routine are provided through a channel
func SomeRoutineChannel[T any](c chan T) {
	defer fmt.Println("Done 3")

	for val := range c {
		fmt.Println(val)
	}
}

// SomeRoutineChannelAndContext - composition of the first and third routine
// It is used in cases when we have in possession no channel
func SomeRoutineChannelAndContext[T any](c <-chan T, ctx context.Context) {
	defer fmt.Println("Done 4")

	for {
		select {
		case val, ok := <-c:
			if ok {
				fmt.Println(val)
			} else {
				return
			}

		case <-ctx.Done():
			return
		}
	}
}

func Task6() {
	// Option 1
	ctx, cansel := context.WithCancel(context.Background())
	go SomeRoutineContext(ctx)
	cansel()

	// Option 2
	var atomicFlag int32
	go SomeRoutineAtomic(&atomicFlag)
	atomic.StoreInt32(&atomicFlag, 1)

	// Option 3
	c := make(chan int)
	go SomeRoutineChannel(c)
	for i := 0; i < 10; i++ {
		if i == 8 {
			// Some event happened
			close(c)
			break
		}

		c <- i
	}

	// Option 4
	// a channel gained from another call
	cc := func() <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()

		return c
	}()
	cctx, ccansel := context.WithCancel(context.Background())
	go SomeRoutineChannelAndContext(cc, cctx)
	ccansel()

	time.Sleep(2 * time.Second)
}
