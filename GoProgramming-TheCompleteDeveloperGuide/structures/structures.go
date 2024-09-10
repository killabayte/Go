package main

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

func area(rect Rectangle)
