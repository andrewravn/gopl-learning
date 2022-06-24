package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Ex 1: modify echo program to print os.Args[0]
// func echo(input []string) {
// 	s, sep := "", ""
// 	for _, arg := range input {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// Ex 2: modify echo program to print index/value of each of its arguments, one per line
// func echo(input []string) {
// 	for i, arg := range input {
// 		fmt.Printf("%v: %v\n", i, arg)
// 	}
// }

// Ex 3: measure runtime difference between our version and strings.join

func track(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s, execution time %s\n", name, time.Since(start))
	}
}

func echo(input []string) {
	for i, arg := range input {
		fmt.Printf("%v: %v\n", i, arg)
	}
}

func echo2(input []string) {
	fmt.Println(strings.Join(input, " "))
}

func main() {
	defer track("echo")()
	echo(os.Args[1:])

	defer track("echo2")()
	echo2(os.Args[1:])
}
