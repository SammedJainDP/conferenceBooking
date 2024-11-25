package validation

import (
	"fmt"
	"strings"
)

func ValidateUser(firstName string, lastName string, email string, remainingTickets uint, userTickets uint) bool {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidticket := userTickets <= remainingTickets && userTickets > 0

	if isValidEmail && isValidName && isValidticket {
		return true
	} else {
		if !isValidEmail {
			fmt.Println("The email entered doesn't contain @. Please Kindly follow the email format")

		}
		if !isValidName {
			fmt.Println("The Name entered should have minimum of 2 charecters. Please Enter the correct First Name and Last Name")

		}

		if !isValidticket {
			fmt.Printf("There are %v tickets left. kindly buy the tickect within the range\n", remainingTickets)

		}

	}
	return false

}
