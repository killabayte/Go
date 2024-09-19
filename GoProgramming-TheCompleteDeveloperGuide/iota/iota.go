package main

type Direction byte

const (
	North = iota
	East
	South
	West
)

north := North
fmt.Println(north)