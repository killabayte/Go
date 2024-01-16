package main

import (
	"fmt"
	"strings"
)

var debugEnabled = true

func main() {
	conferenceName := "Go lang for DevOps Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 {
		var firstName, lastName, email string
		var userTickets uint

		fmt.Println("Please enter your first name, last name and email separated by space:")
		fmt.Scan(&firstName, &lastName, &email)

		fmt.Println("How many tickets you want to book?")
		fmt.Scan(&userTickets)

		var isNameValid = len(firstName) >= 2 && len(lastName) > 2
		var isEmailValid = strings.Contains(email, "@")
		var isTicketsValid = userTickets > 0 && userTickets <= remainingTickets

		if isNameValid && isEmailValid && isTicketsValid {
			remainingTickets -= userTickets

			bookings = append(bookings, firstName+" "+lastName+" "+email+" "+fmt.Sprint(userTickets))

			fmt.Printf("Thanks %s %s, you have booked %d tickets. You will receive a confiramtion on your email: %s\n", firstName, lastName, userTickets, email)
			fmt.Println("Remaining tickets are:", remainingTickets)
			fmt.Printf("All bookings what we have so far: %v\n", bookings)
			if debugEnabled {
				printFirstNames(bookings)
			}
		} else {
			if !isNameValid {
				fmt.Printf("Sorry, your name is not valid. Please try again.\n")
			}
			if !isEmailValid {
				fmt.Printf("Sorry, your email is not valid. Please try again.\n")
			}
			if !isTicketsValid {
				fmt.Printf("Sorry, your tickets number is not valid. Please try again.\n")
			}
		}
	}
}

func greetUsers(c string, t uint, rt uint) {
	fmt.Printf("Wellcome to our %s conference booking application\n", c)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", t, rt)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(b []string) {
	firstNames := []string{}
	for _, booking := range b {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("All first names what we have so far: %v\n", firstNames)
}
