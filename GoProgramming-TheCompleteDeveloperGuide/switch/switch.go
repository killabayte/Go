package main

import "fmt"

func main() {
	switch age := 99; {
	case age == 0:
		fmt.Println("Newborn")
	case age >= 1 && age <= 3:
		fmt.Println("Toddler")
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}
