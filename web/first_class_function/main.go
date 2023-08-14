package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	twice(sayHi)
	twice(sayBye)
	var mathFunction func(int, int) float64
	fmt.Println(mathFunction(5, 2))
}
