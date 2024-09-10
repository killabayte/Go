package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func lenght(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

func area(rect Rectangle) int {
	return (lenght(rect) * width(rect))
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (lenght(rect) * 2)
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {

}
