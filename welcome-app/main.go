package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args
	fmt.Printf("Hi %v! You are pretty mutch wellcome to that awesome Golang for DevOps cource!", name[1:])
}
