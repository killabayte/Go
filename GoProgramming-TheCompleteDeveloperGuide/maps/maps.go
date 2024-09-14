package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 5
	shoppingList["bread"] += 3

	shoppingList["eggs"] = 5
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("After cleanup:", shoppingList)
}
