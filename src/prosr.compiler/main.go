package main

import (
	"fmt"
	"os"
	"os/user"

	"prosr.compiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the prosr.compiler!\n", user.Username)
	fmt.Printf("Feel free to type in definitions :)\n")

	repl.Start(os.Stdin, os.Stdout)
}
