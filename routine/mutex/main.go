package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	counter := 0

	gs := 100

	wg.Add(gs)

	var m sync.Mutex

	for i := 0; i < gs; i++ {

		go func() {
			id := i

			m.Lock()

			counter++

			fmt.Printf("ID: %d, Value: %d\n", id, counter)

			m.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("end value:", counter)
}
