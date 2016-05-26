package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()

	r.GET("/", HomeHandler)

	// posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// posts singualr
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	location := ":8080"

	fmt.Println("Started server on " + location)
	http.ListenAndServe(location, r)
}

func HomeHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

func PostShowHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing single post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "single post update")
}

// XXX: unused
func PostDeleteHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, rq *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
