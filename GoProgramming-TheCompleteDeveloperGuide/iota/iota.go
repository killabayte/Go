package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func main() {
	north := North
	fmt.Println(north)
}
