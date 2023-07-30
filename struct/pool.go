package main

import "fmt"

func main() {
	location := geo.Coordinates{Latitude: 37.42, longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("longitude:", location.longitude)
}
