package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wait sync.WaitGroup

	for i := 0; i < 2; i++ {
		wait.Add(1)

		id := i
		go func() {
			fmt.Println("Print in goroutine: ", id)

			wait.Done()
		}()
	}

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wait.Wait()

	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
