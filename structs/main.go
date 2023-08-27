package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Cloud",
		contactInfo: contactInfo{
			email:   "jim@go.lang",
			zipCode: 10101,
		},
	}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p.firstName)
}
