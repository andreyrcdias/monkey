package main

import (
	"fmt"
	"github.com/andreyrcdias/monkey/repl"
	"os"
)

func main() {
	// TODO: how to proper version a golang package
	version := "0.0.1"
	fmt.Printf("Monkey Language %s\n", version)
	// TODO: implement help, credits commands
	// fmt.Printf("Type \"help\", \"credits\" for more information.\n")
	repl.Start(os.Stdin, os.Stdout)
}
