package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"net/http"
)

func main() {
	location := ":8000"
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Visit `localhost" + location + "`")
	http.ListenAndServe(location, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	formData := []byte(r.FormValue("body"))
	markdown := blackfriday.MarkdownCommon(formData)
	rw.Write(markdown)
}
