//pass_fail print pass user exam or not

package main

import (
	"fmt"
	"log"

	"github.com/killabayte/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A garde of", grade, "is", status)
}
