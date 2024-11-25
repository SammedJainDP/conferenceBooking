package validation

import (
	"fmt"
	"strings"
)

func ValidateUserFirstName(firstName string) bool {

	isValidName := len(firstName) >= 2

	if isValidName {
		return true
	} else {

		fmt.Println("The Name entered should have minimum of 2 charecters. Please Enter the correct First Name.")
		return false
	}

}
func ValidateUserLastName(lastName string) bool {
	isValidName := len(lastName) >= 2
	if isValidName {
		return true
	} else {
		fmt.Println("The Name entered should have minimum of 2 charecters. Please Enter the correct Last Name.~")
		return false
	}

}

func ValidateUserEmail(email string) bool {
	isValidEmail := strings.Contains(email, "@")
	if isValidEmail {
		return true
	} else {
		fmt.Println("The email entered doesn't contain @. Please Kindly follow the email format")
		return false

	}
}

func ValidateUserTickets(remainingTickets uint, userTickets uint) bool {
	isValidticket := userTickets <= remainingTickets && userTickets > 0
	if isValidticket {
		return true
	} else {
		fmt.Printf("Number of tickets entered are wrong. kindly buy the tickect within the range of 1 to  %v \n", remainingTickets)
		return false
	}

}
