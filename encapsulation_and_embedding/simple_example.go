package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

}
