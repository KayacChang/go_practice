package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	inc := 0

	gs := 100

	wait.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := inc

			runtime.Gosched()

			v++

			inc = v

			fmt.Println(inc)
			wait.Done()
		}()
	}

	wait.Wait()

	fmt.Println("end value: ", inc)
}
