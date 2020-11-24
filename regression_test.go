/*
FIXME: I'm having hard time using golang.org/x/tools/go/analysis/analysistest
with the `recheck:<n>` comments.
*/
package recheck

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"regexp"
	"testing"
)

// ./bad.go:20:2: argument 0 not a literal
var errRe = regexp.MustCompile(`\.go:\d+:\d+`)

// FIXME: go.mod & go.test are changed after running the test

func runCmd(t *testing.T, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	return cmd, cmd.Run()
}

func TestOK(t *testing.T) {
	_, err := runCmd(t, "go", "run", "./cmd/recheck", "testdata/ok.go")
	if err != nil {
		t.Fatal(err)
	}
}

func errCount(t *testing.T, r io.Reader) int {
	s := bufio.NewScanner(r)
	n := 0
	for s.Scan() {
		if errRe.FindString(s.Text()) != "" {
			n++
		}
	}

	if err := s.Err(); err != nil {
		t.Fatal(err)
	}
	return n

}

func TestBad(t *testing.T) {
	cmd, err := runCmd(t, "go", "run", "./cmd/recheck", "testdata/bad.go")
	if err == nil {
		t.Fatal(err)
	}

	const numErrors = 5
	buf := cmd.Stderr.(*bytes.Buffer)
	if n := errCount(t, buf); n != numErrors {
		t.Fatalf("expected %d errors, found %d", numErrors, n)
	}
}
