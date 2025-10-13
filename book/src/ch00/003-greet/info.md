El objetivo del programa main.go es generar un saludo, pero para esto vamos a tratar de explicar lo que hemos entendido.

Primero revisemos el codigo  de este programa simple.

```go
package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Mundo", "un nombre para decir hola")
	flag.Parse()
}

func main() {
	if name == "" {
		fmt.Println("Please provide a name using the -name flag.")
		return
	}
	fmt.Printf("Hello, %s!\n", name)
}

```

Lo primero declaramos el nombre del paquete en este caso y como dentro de la carpeta no hay mas archivos go pues este archivo sera el principal por eso es package main.

