package main

import (
	"fmt"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

// a type for a method attached to a controller
type Action func(rw http.ResponseWriter, r *http.Request) error

// a type for the base controller
type AppController struct{}

// the controller we'll use for this mini-app
type MyController struct {
	AppController
	*render.Render
}

// finally, our controller
func (c *AppController) Action(act Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := act(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {

	responseData := map[string]string{"Hello": "JSON"}

	c.JSON(rw, http.StatusOK, responseData)

	return nil
}

func main() {
	lctn := ":8080"
	ctrl := &MyController{Render: render.New(render.Options{})}
	fmt.Printf("Open localhost%s\n", lctn)
	http.ListenAndServe(lctn, ctrl.Action(ctrl.Index))
}
