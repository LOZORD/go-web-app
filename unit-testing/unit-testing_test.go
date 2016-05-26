package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	HelloWorld(res, req)

	expected := "Hello world!"

	actual := res.Body.String()

	if actual != expected {
		fatal_str := fmt.Sprintf("Expected `%s`, but got `%s`", expected, actual)
		t.Fatal(fatal_str)
	}
}
