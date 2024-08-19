package main

import (
	"github.com/andreyrcdias/monkey/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
