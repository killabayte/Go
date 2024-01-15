package main

import "fmt"

func main() {
	var conferenceName = "Go for DevOps Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string

}
