package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	tomas := person{}
	tomas.firstName = "Tomas"
	tomas.lastName = "Anderson"
	fmt.Println(tomas)
}
