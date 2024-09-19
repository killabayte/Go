package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return fmt.Sprintf("North")
	}
}

func main() {
	north := North
	fmt.Println(north)
}
