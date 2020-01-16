package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		fmt.Println(<-c)
	}()

	c <- 42
}
