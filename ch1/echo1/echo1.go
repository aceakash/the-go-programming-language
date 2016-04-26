package main

import (
	"fmt"
	"strings"
	"time"
	"os"
)

func main() {
	myjoinStart := time.Now()
	s1 := myjoin(os.Args[1:])
	myjoinEnd := time.Since(myjoinStart).Nanoseconds()
	fmt.Printf("myjoin: '%s', took %v ns\n", s1, myjoinEnd)

	stringsJoinStart := time.Now()
	s2 := strings.Join(os.Args[1:], " ")
	stringsJoinEnd := time.Since(stringsJoinStart).Nanoseconds()
	fmt.Printf("strings.Join: '%s', took %v ns\n", s2, stringsJoinEnd)
}

func myjoin(arr []string) string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
