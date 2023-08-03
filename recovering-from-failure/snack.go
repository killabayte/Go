package main

import "fmt"

func takoHunt() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}
func snack() {
	defer takoHunt()
	fmt.Println("Opening refrigerator")
	fmt.Println("Closing refrigerator")
	err := fmt.Errorf("There's an error")
	panic(err)
}

func main() {
	snack()
	fmt.Println("no panic! Found a yesterday tako in a chair!")
}
