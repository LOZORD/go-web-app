package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Random struct {
	Value      int
	Randomness string
}

func main() {
	location := ":8080"
	fmt.Println("Go to localhost" + location)
	http.HandleFunc("/", ShowRandom)
	http.ListenAndServe(location, nil)
}

func ShowRandom(w http.ResponseWriter, r *http.Request) {
	rand := Random{9, "Very Random"}
	file_path := path.Join("templates", "index.html")
	tmpl, parse_err := template.ParseFiles(file_path)

	if parse_err != nil {
		http.Error(w, parse_err.Error(), http.StatusInternalServerError)
		return
	}

	exec_err := tmpl.Execute(w, rand)

	if exec_err != nil {
		http.Error(w, exec_err.Error(), http.StatusInternalServerError)
	}
}
