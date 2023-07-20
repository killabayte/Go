//pass_fail print pass user exam or not

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Printf("Read error: %s", err)
	// }
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if input >= 60 {
		status := "passing"
	} else {
		status := "failing"
	}
	fmt.Println(input)
}
