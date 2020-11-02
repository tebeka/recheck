package main

import (
	"net/http"
	"regexp"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z`)
)

func handleFunc(func(http.ResponseWriter, *http.Request), string) {}
func handler(w http.ResponseWriter, r *http.Request)              {}

func main() {
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:1

	// bad type
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:0
}
