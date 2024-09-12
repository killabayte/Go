package main

import "fmt"

type shoppingList struct {
	name  string
	price float64
}

func main() {
	shoppingListActual := [4]shoppingList{
		{name: "apple", price: 15.0},
		{name: "milk", price: 2.5},
		{name: "bekon", price: 22.5},
	}

	lastElement := shoppingListActual[(len(shoppingListActual) - 1)]
	fmt.Println(lastElement)
	totalListNum := len(shoppingListActual)
	fmt.Println(totalListNum)

	var totalCost float64
	for i := 0; i < len(shoppingListActual); i++ {
		totalCost += shoppingListActual[i].price
	}
	fmt.Println(totalCost)
	shoppingListActual[3] = shoppingList{name: "avocado", price: 4.5}o
	fmt.Println(shoppingListActual)
}
