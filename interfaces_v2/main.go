package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func main() {

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hola!"
}
