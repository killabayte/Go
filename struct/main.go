package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var address magazine.Address
	address.Street = "1357 Odd st"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)
}
