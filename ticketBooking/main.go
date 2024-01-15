package main

import "fmt"

func main() {
	conferenceName := "Go lang for DevOps Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings = [50]int{}

	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Please enter your first name, last name and email separated by space:")
	fmt.Scan(&firstName, &lastName, &email)
	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)
	remainingTickets -= userTickets
	fmt.Printf("Thanks %s %s, you have booked %d tickets. You will receive a confiramtion on your email: %s\n", firstName, lastName, userTickets, email)
	fmt.Println("Remaining tickets are:", remainingTickets)

}
