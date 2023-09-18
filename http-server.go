package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Specify routes and handlers
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Start the server
	http.ListenAndServe(":8090", nil)
}
