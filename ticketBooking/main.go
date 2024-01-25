package main

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTickets = 50

var (
	debugEnabled          = false
	conferenceName        = "Go lang for DevOps Conference"
	remainingTickets uint = conferenceTickets
	bookings              = make([]UserData, 0)
)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

// var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isNameValid, isEmailValid, isTicketsValid := validateUserInput(firstName, lastName, email, userTickets)

		if isNameValid && isEmailValid && isTicketsValid {

			bookTicket(userTickets, firstName, lastName, email)
			// wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			if debugEnabled {
				firstNames := getFirstNames()
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
	// wg.Wait()
}

func greetUsers() {
	fmt.Printf("Wellcome to our %s conference booking application\n", conferenceName)
	fmt.Printf("We have total %d tickets for the conference, and %d still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking.firstName)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thanks %s %s, you have booked %d tickets. You will receive a confiramtion on your email: %s\n", firstName, lastName, userTickets, email)
	fmt.Println("Remaining tickets are:", remainingTickets)
	fmt.Printf("All bookings what we have so far: %v\n", bookings)
}

func validateUserInput(fn string, ln string, e string, ut uint) (bool, bool, bool) {
	var isNameValid = len(fn) >= 2 && len(ln) > 2
	var isEmailValid = strings.Contains(e, "@")
	var isTicketsValid = ut > 0 && ut <= remainingTickets
	return isNameValid, isEmailValid, isTicketsValid
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println()
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println()
	//wg.Done()
}
