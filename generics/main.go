package main

import "fmt"

func main() {
	var t1 int = 41
	fmt.Printf("plusOne: %d\n", plusOne(t1))
}

func plusOne(t int) int {
	return t + 1
}
