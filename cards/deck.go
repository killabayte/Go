package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := string{"Ace", "Two", "Three", "Four"}

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
