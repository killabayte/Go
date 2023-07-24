package main

import "fmt"

func main() {
	amount := 6
	//When we use pointers we able to work with variable with diiferent Scope - behind the function for example
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
