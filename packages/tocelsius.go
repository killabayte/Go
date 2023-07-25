package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a tempreture in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
