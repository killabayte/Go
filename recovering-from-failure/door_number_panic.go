package main

import (
	"fmt"
	"math/rand"
)

func awardPrize() {
	doorNumber := rand.Intn(4) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a rock!")
	} else {
		panic("Invalid dooe number")
	}
}

func main() {
	awardPrize()
}
