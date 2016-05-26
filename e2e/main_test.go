package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())

	defer ts.Close()

	res, err := http.Get(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	part1 := "Before calling `next`..."
	part2 := "Hello world!"
	part3 := "...After calling `next`"

	expected := part1 + part2 + part3

	actual := string(body)

	if actual != expected {
		fatal_str := fmt.Sprintf("Expected `%s`, but got `%s`", expected, actual)
		t.Fatal(fatal_str)
	}
}
