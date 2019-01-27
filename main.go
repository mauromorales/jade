package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/mauromorales/jade/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hola %s! Esta es la consola interactiva de Jade\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
