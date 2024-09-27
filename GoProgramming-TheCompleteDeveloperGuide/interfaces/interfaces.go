package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add dressing")

}

func main() {

}
