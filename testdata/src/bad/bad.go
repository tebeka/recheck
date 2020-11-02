package bad

import (
	"net/http"
	"regexp"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z`)      // want `missing closing ] in "[a-zA-Z\""`
	re2 = regexp.MustCompilePOSIX(`[a-zA-Z`) // want `missing closing ] in "[a-zA-Z\""`
)

func handleFunc(func(http.ResponseWriter, *http.Request), string) {}
func handler(w http.ResponseWriter, r *http.Request)              {}
func inc(n int) int                                               { return n + 1 }

func init() {
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:1 want `missing closing ] in "[0-9+}\""`

	// not a literal
	handleFunc(handler, "/users/{id:[0-9+}") // recheck:0 want `argument 0 not a literal`

	// not a string
	inc(3) // recheck:0 want `argument 0 not a string`
}
