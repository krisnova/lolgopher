package main

import (
	"github.com/CrowdSurge/banner"
	"github.com/kris-nova/lolgopher/lol"
)

func main() {
	w := lol.NewLolWriter()
	w.Write([]byte(banner.PrintS("lolgopher")))
}
