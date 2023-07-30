package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	address := magazine.Address{Street: "1357 Odd St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)

}
