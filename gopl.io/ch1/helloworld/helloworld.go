package main

import (
	"fmt"
	"os"
)

func main() {

	// example iterating over cli input
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Hello, %v\n", os.Args[i])
	}
}
