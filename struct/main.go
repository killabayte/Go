package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "1357 Odd St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Subscriber name:", subscriber.Name)
	fmt.Println("Subscriber street:", subscriber.Street)
	fmt.Println("Subscriber city:", subscriber.City)
	fmt.Println("Subscriber state:", subscriber.State)
	fmt.Println("Subscriber postal code:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "2468 Even St"
	employee.City = "portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee name:", employee.Name)
	fmt.Println("Employee street:", employee.Street)
	fmt.Println("Employee city:", employee.City)
	fmt.Println("Employee state:", employee.State)
	fmt.Println("Employee postal code:", employee.PostalCode)

}
