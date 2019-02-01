package main

import (
	"fmt"
	"os"

	"github.com/mauromorales/jade/repl"
)

func main() {
	fmt.Println("Bienvenido a la consola interactiva de Jade")
	fmt.Println("Para salir presiona ^C (Control-C)")
	repl.Start(os.Stdin, os.Stdout)
}
