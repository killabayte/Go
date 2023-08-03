package main

import "fmt"

func takoHunt() {
	fmt.Println(recover())
}
func snack() {
	defer takoHunt()
	fmt.Println("Opening refrigerator")
	fmt.Println("Closing refrigerator")
	panic("refrigerator is empty")
}

func main() {
	snack()
	fmt.Println("Found a yesterday tako in a chair, no panic!")
}
