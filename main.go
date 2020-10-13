package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%s=%s\n", k, v)
	}
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}
