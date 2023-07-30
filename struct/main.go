package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress.Street = "1357 Odd St"
	subscriber.HomeAddress.City = "Omaha"
	subscriber.HomeAddress.State = "NE"
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber name:", subscriber.Name)
	fmt.Println("Subscriber street:", subscriber.HomeAddress.Street)
	fmt.Println("Subscriber city:", subscriber.HomeAddress.City)
	fmt.Println("Subscriber state:", subscriber.HomeAddress.State)
	fmt.Println("Subscriber postal code:", subscriber.HomeAddress.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.HomeAddress.Street = "2468 Even St"
	employee.HomeAddress.City = "portland"
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println("Employee name:", employee.Name)
	fmt.Println("Employee street:", employee.HomeAddress.Street)
	fmt.Println("Employee city:", employee.HomeAddress.City)
	fmt.Println("Employee state:", employee.HomeAddress.State)
	fmt.Println("Employee postal code:", employee.HomeAddress.PostalCode)

}
