package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/tebeka/recheck"
)

func main() {
	singlechecker.Main(recheck.Analyzer)
}
