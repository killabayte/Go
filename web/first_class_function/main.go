package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func sayBye() {
	fmt.Println("Bye")
}

func main() {
	var myFunction func()
	myFunction = sayHi
	myFunction()
}
