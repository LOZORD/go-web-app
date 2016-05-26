package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	location := ":8080"
	fmt.Println("Go to localhost" + location)
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(location, nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{Title: "Catch-22", Author: "Joseph Heller"}

	json_result, err := json.Marshal(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Sending: " + string(json_result))

	w.Header().Set("Content-Type", "application/json")
	w.Write(json_result)
}
