package main

type englishBot struct {
}

type spanishBot struct {
}

func main() {

}

func (eb englishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Very custom logic for generation an greeting
	return "Hola!"
}
