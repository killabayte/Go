package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go lang for DevOps Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
		var firstName, lastName, email string
		var userTickets uint

		fmt.Println("Please enter your first name, last name and email separated by space:")
		fmt.Scan(&firstName, &lastName, &email)

		fmt.Println("How many tickets you want to book?")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets

			bookings = append(bookings, firstName+" "+lastName+" "+email+" "+fmt.Sprint(userTickets))

			fmt.Printf("Thanks %s %s, you have booked %d tickets. You will receive a confiramtion on your email: %s\n", firstName, lastName, userTickets, email)
			fmt.Println("Remaining tickets are:", remainingTickets)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("All bookings what we have so far: %v\n", bookings)
			fmt.Printf("All first names what we have so far: %v\n", firstNames)
		} else {
			fmt.Printf("Sorry, we only have %d tickets left. Please try again with less tickets.\n", remainingTickets)
			continue
		}
	}
}
