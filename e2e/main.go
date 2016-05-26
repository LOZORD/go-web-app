package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(res, "Hello world!")
}

func App() http.Handler {
	n := negroni.Classic()

	m := func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		fmt.Fprint(res, "Before calling `next`...")
		next(res, req)
		fmt.Fprint(res, "...After calling `next`")
	}

	n.Use(negroni.HandlerFunc(m))

	r := httprouter.New()

	r.GET("/", HelloWorld)

	n.UseHandler(r)

	return n
}

func main() {
	http.ListenAndServe(":3000", App())
}
