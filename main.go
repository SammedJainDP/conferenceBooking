package main

import (
	"bufio"
	"conferenceTicketBooking/cities"
	"conferenceTicketBooking/user"
	"conferenceTicketBooking/validation"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "conferenceDB")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	var firstName string
	var lastName string
	var email string
	var remainingTickets uint
	var cityId int
	var userTickets uint
	var flag bool = true
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("welcome to the Conference Ticket Boolking CLI Application")

	for {

		if flag {
			remainingTickets, cityId = cities.GetCities(db)
		}

		fmt.Println("Kindly fill the details below to avail the tickets")
		for {
			fmt.Println("Enter the First Name:")
			// fmt.Scanln(&firstName)
			firstName, err = reader.ReadString('\n')
			if err != nil {
				log.Panic("could not read the first name", err)
			}

			if validation.ValidateUserFirstName(firstName) {
				break
			}
			continue
		}
		for {

			fmt.Println("Enter the Last Name:")
			lastName, err = reader.ReadString('\n')
			if err != nil {
				log.Panic("could not read the last name", err)
			}

			if validation.ValidateUserLastName(lastName) {
				break
			}
			continue
		}

		for {
			fmt.Println("Enter the email:")
			fmt.Scan(&email)
			if validation.ValidateUserEmail(email) {
				break
			}
			continue
		}
		for {
			fmt.Println("Enter the number of tickets you want to purchase :")
			fmt.Scan(&userTickets)
			if validation.ValidateUserTickets(remainingTickets, userTickets) {
				break
			}
			continue
		}
		fmt.Printf("thank you %v %v for booking the tickets. You will get the confirmation shortly to %v\n", firstName, lastName, email)
		user.UpdateUserDetails(db, firstName, lastName, email, userTickets, cityId)
		remainingTickets = remainingTickets - userTickets
		cities.UpdateRemainingTickets(db, remainingTickets, cityId)
		flag = true
		sendTickets(firstName, email, userTickets)
		continue
	}
}

func sendTickets(firstName string, email string, userTickets uint) {
	var ticket = fmt.Sprintf("Hi %v,\n you have booked %v tickets.\n", firstName, userTickets)
	fmt.Println("###############")
	fmt.Printf("%v The ticket confirmation has been sent to the Email: %v\n ", ticket, email)
	fmt.Println("###############")
}
