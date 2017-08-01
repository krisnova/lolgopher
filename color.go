package lol

type colorer interface {
	rainbow(freq float64, i float64)
	format() []byte
	reset() []byte
}
