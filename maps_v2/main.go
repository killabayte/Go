package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#fF0000",
		"green": "#00FF00",
		"blue":  "#bbbbbb",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color: %v, hex: %v", color, hex)
	}
}
