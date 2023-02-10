package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Args

	if len(name) < 2 {
		fmt.Println("Usege: ./main <name>")
		os.Exit(1)
	}

	fmt.Printf("Hi %v, you are pretty much wellcome to this Golang cource for DevOps", name[1:])
}
