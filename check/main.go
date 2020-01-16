package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime: %s\n", runtime.GOOS)

	fmt.Printf("architecture: %s\n", runtime.GOARCH)
}
