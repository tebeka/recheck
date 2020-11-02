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
func inc(n int) int                                               { return n + 1 }

func main() {
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:1

	// not a literal
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:0

	// not a string
	inc(3) // recheck:0
}
