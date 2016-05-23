package main

import (
	"fmt"
	"net/http"
)

func main() {
	location := ":8000"
	fmt.Println("Visit `localhost" + location + "`")
	http.ListenAndServe(location, http.FileServer(http.Dir(".")))
}
