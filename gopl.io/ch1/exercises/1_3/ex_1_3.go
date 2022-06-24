package main

// Example usage:
// andrewravn@andrews-mbp-2 exercises % echo -n "test\ntest\nother\n" | go run ex_1_3.go
// 2	test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Example usage:
// andrewravn@andrews-mbp-2 exercises % echo -n "test\ntest\nother\n" | go run ex_1_3.go
// 2	test
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Example usage:
// andrewravn@andrews-mbp-2 exercises % go run ex_1_3.go test_file
// 2	test
// 2	as
func dup2() {
	counts := make(map[string]int)
	fileTracker := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "", counts, fileTracker)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, fileTracker)
			f.Close()
		}
	}

	for file, l := range fileTracker {
		fmt.Println("Filename: ", file)
		for line, n := range counts {
			if strings.Contains(l, line) {
				fmt.Printf("\t%s\t%d\n", line, n)
			}
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]int, fileTracker map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !strings.Contains(fileTracker[fileName], input.Text()) {
			fileTracker[fileName] = fileTracker[fileName] + input.Text() + ","
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	// dup1()
	dup2()
	// dup3()
}
