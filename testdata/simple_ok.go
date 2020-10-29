package main

import (
	"regexp"
)

var (
	re1 = regexp.MustCompile(`[a-zA-Z]+`)
	re2 = regexp.MustCompile(`\d+`)
	re3 = regexp.MustCompile(`\d`)
)
