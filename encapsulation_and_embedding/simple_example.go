package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) {
	d.Month = month
}

func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	date := Date{}
	err := date.SetYear(-2019)
	if err != nil {
		log.Fatal(err)
	}
	date.SetMonth(70)
	date.SetDay(42)
	fmt.Println(date)
}
