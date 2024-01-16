package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var (
	debugEnabled          = false
	conferenceName        = "Go lang for DevOps Conference"
	remainingTickets uint = conferenceTickets
	bookings              = []string{}
)

func main() {
	greetUsers()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isNameValid, isEmailValid, isTicketsValid := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isNameValid && isEmailValid && isTicketsValid {
			bookTicket(remainingTickets, userTickets, firstName, lastName, email, bookings)
			if debugEnabled {
				firstNames := getFirstNames(bookings)
				fmt.Printf("All first names what we have so far: %v\n", firstNames)
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

func greetUsers() {
	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(b []string) []string {
	firstNames := []string{}
	for _, booking := range b {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(fn string, ln string, e string, ut uint, rt uint) (bool, bool, bool) {
	var isNameValid = len(fn) >= 2 && len(ln) > 2
	var isEmailValid = strings.Contains(e, "@")
	var isTicketsValid = ut > 0 && ut <= rt
	return isNameValid, isEmailValid, isTicketsValid
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Please enter your first name, last name and email separated by space:")
	fmt.Scan(&firstName, &lastName, &email)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, bookings []string) {
	remainingTickets -= userTickets

	bookings = append(bookings, firstName+" "+lastName+" "+email+" "+fmt.Sprint(userTickets))

	fmt.Printf("Thanks %s %s, you have booked %d tickets. You will receive a confiramtion on your email: %s\n", firstName, lastName, userTickets, email)
	fmt.Println("Remaining tickets are:", remainingTickets)
	fmt.Printf("All bookings what we have so far: %v\n", bookings)
}
