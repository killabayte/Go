package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var kia car
	kia.name = "rio"
	kia.topSpeed = 220
	fmt.Println("Name:", kia.name)
	fmt.println("Top speed:", kia.topSpeed)

	var bolts parts
	bolts.description = "hex bolts"
	bolts.count = 42
	fmt.Println("Description:", bolts.description)
	fmt.println("Count:", bolts.count)
}
