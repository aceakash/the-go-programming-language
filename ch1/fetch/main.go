package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://untilfalse.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	fmt.Println(string(b))
}
