package main

import "fmt"

func main() {
	var myName = "John"
	fmt.Println("My name is:", myName)

	var name string = "Kate"
	fmt.Println("Name is equal:", name)

	userName := "admin"
	fmt.Println("User name is:", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("The sum of part1 and other is", part1+other)

	part2, other := 2, 0
	fmt.Println("part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("New sum is:", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}
