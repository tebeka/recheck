package recheck

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestBad1(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "bad")
}
