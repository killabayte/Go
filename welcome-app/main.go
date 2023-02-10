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

	fmt.Printf("Hi %v! You are pretty mutch wellcome to that awesome Golang for DevOps cource!", name[1:])
}
