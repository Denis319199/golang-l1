package task5

import (
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context) {
	c := make(chan int)

	go func() {
		defer close(c)

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case c <- i:
			}
		}
	}()

	for val := range c {
		fmt.Println(val)
	}
}

func Task5() {
	ctx, cansel := context.WithTimeout(context.Background(), 2*time.Second)
	Run(ctx)
	cansel()
}
