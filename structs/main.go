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
	jim := person{
		firstName: "Jim",
		lastName:  "Cloud",
		contact: contactInfo{
			email:   "jim@go.lang",
			zipCode: 10101,
		},
	}
	fmt.Printf("%+v", jim)
}
