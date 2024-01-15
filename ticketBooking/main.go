package main

import "fmt"

func main() {
	conferenceName := "Go for DevOps Conference"
	const conferenceTickets = 50
	remainingTickets := conferenceTickets

	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName, lastName, email string
	var userTickets int

	fmt.Println("Please enter your first name, last name and email separated by space:")
	fmt.Scan(&firstName, &lastName, &email)
	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)
	fmt.Printf("User %s %s booked %d tickets\n", firstName, lastName, userTickets)

}
