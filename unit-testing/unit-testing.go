package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello world!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":3000", nil)
}
