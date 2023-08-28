package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#fF0000",
	// 	"green": "#00FF00",
	// }
	colors := make(map[string]string)
	colors["blue"] = "#0000FF"
	fmt.Println(colors)
}
