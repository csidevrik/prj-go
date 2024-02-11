package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	// Crear instancias de colores
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	// Imprimir mensajes en diferentes colores
	fmt.Println("Mensaje en color:", red("rojo"))
	fmt.Println("Mensaje en color:", green("verde"))
	fmt.Println("Mensaje en color:", yellow("amarillo"))
	fmt.Println("Mensaje en color:", blue("azul"))

	// Puedes combinar colores
	fmt.Println("Mensaje en color combinado:", red("rojo"), blue("azul"))

	// Salir con estado de éxito
	os.Exit(0)
}
