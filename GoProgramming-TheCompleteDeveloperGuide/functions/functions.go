package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(l, r int) int {
	return l + r
}

func greet() {
	fmt.Println("Hi, this is great day!")
}

func main() {
	greet()
	dozen := double(6)
	fmt.Println("This is dozen:", dozen)
	bakerDozen := add(dozen, 1)
	fmt.Println("A baker dozen is:", bakerDozen)
	anotherBakerDozen := add(double(6), 1)
	fmt.Println("Another bakers Dozen is:", anotherBakerDozen)
}
