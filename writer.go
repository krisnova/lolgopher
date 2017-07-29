package lol

import (
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	spread = 3.0
	freq   = 0.1
	seed   = 0
)

type color struct {
	R, G, B uint8
}

func (c *color) rainbow(freq float64, i float64) {
	c.R = uint8(math.Floor(math.Sin(freq*i+0)*127)) + 128
	c.G = uint8(math.Floor(math.Sin(freq*i+2.0*math.Pi/3.0)*127)) + 128
	c.B = uint8(math.Floor(math.Sin(freq*i+4.0*math.Pi/3.0)*127)) + 128
}

func (c *color) format() []byte {
	return []byte(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B))
}

// LolWriter writes a little lol-er.
type Writer struct {
	Output  io.Writer
	lineIdx int
	origin  int
}

func (w *Writer) Write(p []byte) (int, error) {
	nWritten := 0
	ss := strings.Split(string(p), "\n")
	for i, s := range ss {
		// TODO: strip out pre-existing ANSI codes and expand tabs. Would be
		// great to expand tabs in a context aware way (line linux expand
		// command).

		n, err := w.writeRaw(s)
		if err != nil {
			return nWritten, err
		}
		nWritten += n

		// Increment the origin (line count) for each newline.  There is a
		// newline for every item in this array except the last one.
		if i != len(ss)-1 {
			n, err := w.Output.Write([]byte("\n"))
			if err != nil {
				return nWritten, err
			}
			nWritten += n
			w.origin++
			w.lineIdx = 0
		}
	}
	return nWritten, nil
}

// writeRaw will write a lol'd s to the underlying writer.  It does no line
// detection.
func (w *Writer) writeRaw(s string) (int, error) {
	var c color
	nWritten := 0
	for _, r := range s {
		c.rainbow(freq, float64(w.origin)+float64(w.lineIdx)/spread)
		_, err := w.Output.Write(c.format())
		if err != nil {
			return nWritten, err
		}
		n, err := w.Output.Write([]byte(string(r)))
		if err != nil {
			return nWritten, err
		}
		nWritten += n
		w.lineIdx++
	}
	return nWritten, nil
}
