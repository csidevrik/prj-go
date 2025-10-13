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
