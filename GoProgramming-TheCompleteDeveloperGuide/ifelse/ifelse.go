package main

import "fmt"

func avarage(a, b, c int) float64 {
	return float64(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 5
	if quiz1 > quiz2 {
		fmt.Println("Quiz1 is greater than Quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz2 is greater than Quiz1")
	} else {
		fmt.Println("Quiz1 and Quiz2 are equal")
	}

	if avarage(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("The avarage score is acceptable")
	} else {
		fmt.Println("Please try again, and re-cap some quiz")
	}
}
