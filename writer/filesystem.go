package writer

import (
	"io"
	"os"
)

type FilesystemLolcatWriter struct {
}

func NewFilesystemLolcatWriter() io.Writer {
	return &StdoutLolgopherWriter{}
}

func (w *FilesystemLolcatWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
