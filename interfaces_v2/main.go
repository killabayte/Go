package main

import "fmt"

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

}

func printEnglishGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSpanishGreeting(sb englishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hola!"
}
