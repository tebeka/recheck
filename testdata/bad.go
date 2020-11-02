package main

import (
	"net/http"
	"regexp"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z`)
)

func handler(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/articles/{category}/{id:[0-9+}", handler) // recheck:0
}
