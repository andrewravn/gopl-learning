package main

import "fmt"

func Repeat(x string, count int) string {
	var res string
	for i := 0; i < count; i++ {
		res += x
	}
	return res
}

func main() {
	fmt.Printf(Repeat("a", 5))
}
