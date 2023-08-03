package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}
func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}
func (r Refrigerator) Findfood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}

func main() {
	frige := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := frige.Findfood(food)
		if err != nil {
			log.Fatal(err)
		}
	}

}
