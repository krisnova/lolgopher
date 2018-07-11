package main

import (
	"bufio"
	"os"

	"github.com/kris-nova/lolgopher/lol"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lol.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
