package main

import (
	"github.com/codegangsta/negroni"
	"log"
	"net/http"
)

const password string = "secret123"

func main() {
	// middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	location := ":8080"

	n.Run(location)
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	req_pw := r.URL.Query().Get("password")

	if req_pw == password {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back")
}
