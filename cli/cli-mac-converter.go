package main

import (
	"flag"
	"fmt"
)

func main() {
	// Definir y parsear banderas
	namePtr 	:= flag.String("name", "Usuario", "Nombre del usuario")
	agePtr 		:= flag.Int("age", 25, "Edad del usuario")
	emailPtr 	:= flag.String("email", "usuario@example.com", "Correo electrónico del usuario")

	flag.Parse()

	// Imprimir la información del usuario
	fmt.Printf("Hola, %s. Edad: %d, Correo electrónico: %s\n", *namePtr, *agePtr, *emailPtr)
}