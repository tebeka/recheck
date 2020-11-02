package main

import (
	"net/http"
	"regexp"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z]+`)
	re2 = regexp.MustCompile(`\d+`)
	re3 = regexp.MustCompile(`\d`)
	re4 = regexp.MustCompilePOSIX(`\d`)
)

func handleFunc(func(http.ResponseWriter, *http.Request), string) {}
func handler(w http.ResponseWriter, r *http.Request)              {}

func main() {
	handleFunc(handler, "/users/{id:[0-9]+}") // recheck:1
}
