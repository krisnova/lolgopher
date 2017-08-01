package main

import (
	"fmt"
	lol "github.com/kris-nova/lolgopher"
)

func main() {
	writer := lol.NewLolWriter()
	fmt.Fprint(writer, "----------------------------------------------------\n")
}
