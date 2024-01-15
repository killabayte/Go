package main

import "fmt"

func main() {
	var conferenceName = "Go for DevOps Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Wellcome to our %s conference booking application", conferenceName)
	fmt.Println("Get your tickets here to attend")

}
