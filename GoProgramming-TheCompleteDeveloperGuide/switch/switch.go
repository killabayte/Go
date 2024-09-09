package main

import "fmt"

func main() {
	switch age := 12; {
	case age == 0:
		fmt.Println("Newborn")
	case age > 0 && age <= 3:
		fmt.Println("Toddler")
	case age > 3 && age <= 12:
		fmt.Println("child")
	case age > 12 && age <= 17:
		fmt.Println("teenager")
	case age >= 18:
		fmt.Println("adult")
	}
}
