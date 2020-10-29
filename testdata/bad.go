package main

import (
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z`)
)

func handler(w http.ResponseWriter, r *http.Request) {}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9+}", handler) // recheck:0
}
