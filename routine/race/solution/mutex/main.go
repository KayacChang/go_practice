package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	var m sync.Mutex

	inc := 0

	gs := 100

	wait.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()

			v := inc

			v++

			inc = v

			fmt.Println(inc)

			m.Unlock()

			wait.Done()
		}()
	}

	wait.Wait()

	fmt.Println("end value: ", inc)
}
