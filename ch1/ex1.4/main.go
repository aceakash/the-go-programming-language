package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNamesByLine := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNamesByLine)
	} else {
		countLinesForFiles(files, counts, fileNamesByLine)
	}

	printOutput(counts, fileNamesByLine)
}

func countLinesForFiles(fileNames []string, counts map[string]int, filesByLine map[string][]string) {
	for _, arg := range fileNames {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, filesByLine)
		f.Close()
	}
}

func printOutput(counts map[string]int, fileNamesByLine map[string][]string) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%q\n", n, line, strings.Join(fileNamesByLine[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNamesByLine map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !contains(fileNamesByLine[line], f.Name()) {
			fileNamesByLine[line] = append(fileNamesByLine[line], f.Name())
		}

	}
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
