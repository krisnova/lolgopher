package writer

import (
	"io"
	"os"
)

type StdoutLolgopherWriter struct {
}

func NewStdoutLolgopherWriter() io.Writer {
	return &StdoutLolgopherWriter{}
}

func (w *StdoutLolgopherWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
