package main

import (
	"fmt"
	"net/http"
)


func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("abb","ccc")

	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hello", hello)


	// serve
	http.ListenAndServe("127.0.0.1:8090", nil)
}
