package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}
}

func countLines(reader io.Reader) {
	counts := make(map[string]int)

	input := bufio.NewScanner(reader)
	for input.Scan() {
		line := input.Text()
		// fmt.Printf("You said: %q\n", line)
		counts[line] = counts[line] + 1
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line, n)
		}
	}
}

/* notes

- using io.Reader instead of *os.File as parameter

- iter 1: remove dups from anywhere
- iter 2: remove only adjacent dups
- iter 3: command line flag to remove only adjacent

*/
