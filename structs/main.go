package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	tomas := person{}
	tomas.firstName = "Tomas"
	tomas.lastName = "Anderson"
	fmt.Println(tomas)
}
