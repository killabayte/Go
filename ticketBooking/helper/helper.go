package helper

import "strings"

func validateUserInput(fn string, ln string, e string, ut uint) (bool, bool, bool) {
	var isNameValid = len(fn) >= 2 && len(ln) > 2
	var isEmailValid = strings.Contains(e, "@")
	var isTicketsValid = ut > 0 && ut <= remainingTickets
	return isNameValid, isEmailValid, isTicketsValid
}
