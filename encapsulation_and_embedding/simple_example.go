package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(07)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

}
