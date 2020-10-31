package recheck

import (
	"os/exec"
	"testing"
)

// FIXME: go.mod & go.test are changed after running the test

func TestOK(t *testing.T) {
	cmd := exec.Command("go", "run", "./cmd/recheck", "testdata/ok.go")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestBad(t *testing.T) {
	cmd := exec.Command("go", "run", "./cmd/recheck", "testdata/bad.go")
	if err := cmd.Run(); err == nil {
		t.Fatal(err)
	}
}
