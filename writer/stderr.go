package writer

import (
	"io"
	"os"
)

type StderrLolgopherWriter struct {
}

func NewStderrLolgopherWriter() io.Writer {
	return &StdoutLolgopherWriter{}
}

func (w *StderrLolgopherWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
