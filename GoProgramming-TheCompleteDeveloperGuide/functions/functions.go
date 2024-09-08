package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(l, r int) int {
	return l + r
}

func greet(name string) {
	fmt.Println("Namaskar", name)
}

func namaste() string {
	return "Namaste!"
}

func trikonasana(a, b, c int) int {
	return a + b + c
}

func numReturn() int {
	return 42
}

func multiNumReturn() (int, int) {
	return 54, 54
}

func addThreeNum(a, b, c int) int {
	return a + b + c
}

func main() {
	greet("Matsyendranatha!")
	fmt.Println(namaste())
	trikonasanaResult := trikonasana(1, 100, 7)
	fmt.Println("Trikonasana result is:", trikonasanaResult)
	b, c := multiNumReturn()
	sumOfThree := addThreeNum(numReturn(), b, c)
	fmt.Println("This is sum of three num:", sumOfThree)
	dozen := double(6)
	fmt.Println("This is dozen:", dozen)
	bakerDozen := add(dozen, 1)
	fmt.Println("A baker dozen is:", bakerDozen)
	anotherBakerDozen := add(double(6), 1)
	fmt.Println("Another bakers Dozen is:", anotherBakerDozen)
}

//test
