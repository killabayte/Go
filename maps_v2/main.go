package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#fF0000",
	// 	"green": "#00FF00",
	// }
	colors := make(map[int]string)
	colors[0] = "zero"
	fmt.Println(colors)
}
