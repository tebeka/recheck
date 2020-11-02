package recheck

import (
	"os"
	"os/exec"
	"testing"
)

// FIXME: go.mod & go.test are changed after running the test

type TestWriter struct {
	prefix string
	t      *testing.T
}

func (t *TestWriter) Write(data []byte) (int, error) {
	t.t.Logf("%s: %s\n", t.prefix, string(data))
	return len(data), nil
}

func runCmd(t *testing.T, args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	if os.Getenv("DEBUG") != "" {
		cmd.Stdout = &TestWriter{"STDOUT", t}
		cmd.Stderr = &TestWriter{"STDERR", t}
	}
	return cmd.Run()
}

func TestOK(t *testing.T) {
	err := runCmd(t, "go", "run", "./cmd/recheck", "testdata/ok.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBad(t *testing.T) {
	err := runCmd(t, "go", "run", "./cmd/recheck", "testdata/bad.go")
	if err == nil {
		t.Fatal(err)
	}
}
