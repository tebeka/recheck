package recheck

import (
	"os/exec"
	"testing"
)

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
