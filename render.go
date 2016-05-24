package main

import (
	"fmt"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		ip_and_port := req.RemoteAddr
		r.HTML(w, http.StatusOK, "example", ip_and_port)
	})

	runFunc(func() {
		fmt.Println("hello")
	})

	// test to check whether &func == func (it does NOT)
	//runFunc(&printHello)

	runFunc(printHello)

	// open `localhost:8080`
	http.ListenAndServe(":8080", mux)
}

func runFunc(someFunc func()) {
	someFunc()
}

func printHello() {
	fmt.Println("printing hello")
	fmt.Printf("type of `printHello`: %T", printHello)
}
