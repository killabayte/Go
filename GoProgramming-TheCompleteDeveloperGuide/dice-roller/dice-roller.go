package main

import (
	"fmt"
	"math/rand"
)

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func diceDrop() (int, int) {
	d1 := randRange(1, 6)
	d2 := randRange(1, 6)
	return d1, d2
}

func main() {

	diceOne, diceTwo := diceDrop()
	diceSum := diceOne + diceTwo

	switch {
	case diceOne == 1 && diceTwo == 1 && diceSum == 2:
		fmt.Println("Snake eyes")
	case diceSum == 7:
		fmt.Println("Lucky 7")
	case diceSum%2 == 0:
		fmt.Println("Even")
	case diceSum%2 == 1:
		fmt.Println("Odd")
	}
	fmt.Println("First dice is:", diceOne)
	fmt.Println("Second dice is:", diceTwo)

}
