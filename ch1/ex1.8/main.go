package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	goFetch(os.Args[1:])
}

func goFetch(urls []string) {
	for _, url := range urls {
		fetchAndDisplayForSingleURL(url)
	}
}

func fetchAndDisplayForSingleURL(url string) {
	resp, err := http.Get(ensureHTTP(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

func ensureHTTP(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}
