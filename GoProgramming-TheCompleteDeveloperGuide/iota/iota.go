package main

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

north := North
fmt.Println(north)