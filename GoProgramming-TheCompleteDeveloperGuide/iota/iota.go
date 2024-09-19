package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

// func (d Direction) String() string {
// 	switch d {
// 	case North:
// 		return fmt.Sprintf("North")
// 	case South:
// 		return fmt.Sprintf("Soutch")
// 	case East:
// 		return fmt.Sprintf("East")
// 	case West:
// 		return fmt.Sprintf("West")
// 	default:
// 		return "other direction"
// 	}
// }

func main() {
	north := North
	fmt.Println(north)
}
