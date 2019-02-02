package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mauromorales/jade/repl"
	"github.com/mauromorales/jade/runner"
)

func main() {
	interactive := flag.Bool("i", false, "Ejecuta la consola interactiva de Jade")
	ayuda := flag.Bool("a", false, "Muestra el mensaje de ayuda de este ejecutable")

	flag.Usage = func() {
		fmt.Printf("Uso de %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *ayuda {
		flag.Usage()
		os.Exit(0)
	}

	if *interactive || len(os.Args) == 1 {
		fmt.Println("Bienvenido a la consola interactiva de Jade")
		fmt.Println("Para salir presiona ^C (Control-C)")
		repl.Start(os.Stdin, os.Stdout)
	}

	file := os.Args[1]

	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("No puedo encontrar el archivo: %s\n", file)
		os.Exit(1)
	}

	dat, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error al leer el archivo: %s\n", file)
	}
	defer dat.Close()

	runner.Start(dat, os.Stderr)
}
