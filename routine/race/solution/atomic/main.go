package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wait sync.WaitGroup

	var inc int64 = 0

	gs := 100

	wait.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)

			fmt.Println(atomic.LoadInt64(&inc))

			wait.Done()
		}()
	}

	wait.Wait()

	fmt.Println("end value: ", inc)
}
