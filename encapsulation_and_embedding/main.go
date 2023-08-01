package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/geo"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Longitude)
	fmt.Println(coordinates.Latitude)
}
