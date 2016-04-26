package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		countDupsInFile(filename)
	}
}

func countDupsInFile(filename string) {
	fmt.Printf("\nIn file %s...\n", filename)

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	counts := make(map[string]int)

	for _, line := range strings.Split(string(buf), "\n") {
		counts[line]++
	}

	for line, count := range counts {
		if count > 1 {
			plural := "s"
			if count == 2 {
				plural = ""
			}
			fmt.Printf("%q was duplicated %d time%s\n", line, count-1, plural)
		}
	}
}
