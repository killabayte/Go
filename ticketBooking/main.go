package main

import "fmt"

func main() {
	conferenceName := "Go for DevOps Conference"
	const conferenceTickets = 50
	remainingTickets := conferenceTickets

	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	userName := "John Wick"
	userTickets := 3
	fmt.Printf("User %s booked %d tickets\n", userName, userTickets)

}
