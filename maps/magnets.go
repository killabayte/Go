package main

import "fmt"

func main() {
	ranks := map[string]int{"gold": 1, "bronze": 3, "silver": 2}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}
}
