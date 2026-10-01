package app

import (
	"bytes"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	var output bytes.Buffer
	if err := Run(&output); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	want := "2 + 3 = 5\n" +
		"division by zero was correctly detected: calculator: divide 10 by 0: calculator: division by zero\n" +
		"word count: 4\n"
	if output.String() != want {
		t.Errorf("Run() output = %q, want %q", output.String(), want)
	}
}

type failingWriter struct {
	err error
}

func (w failingWriter) Write([]byte) (int, error) {
	return 0, w.err
}

func TestRunReturnsOutputError(t *testing.T) {
	writeErr := errors.New("write failed")
	err := Run(failingWriter{err: writeErr})
	if !errors.Is(err, writeErr) {
		t.Fatalf("Run() error = %v, want wrapped %v", err, writeErr)
	}
}
