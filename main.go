package main

import (
	"conferenceTicketBooking/cities"
	"conferenceTicketBooking/user"
	"conferenceTicketBooking/validation"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "modernc.org/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

var waitGroup = sync.WaitGroup{}

func main() {
	db, err := sql.Open("sqlite3", "conferenceDB")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()
	// createTable(db)
	var firstName string
	var lastName string
	var email string
	var remainingTickets uint
	var cityId int
	var userTickets uint

	var flag bool = true
	fmt.Println("welcome to the Conference Ticket Boolking CLI Application")

	for {

		if flag {
			remainingTickets, cityId = cities.GetCities(db)
		}

		fmt.Println("Kindly fill the details below to avail the tickets")

		fmt.Println("Enter the First Name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter the Last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter the email:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to purchase :")
		fmt.Scan(&userTickets)

		val := validation.ValidateUser(firstName, lastName, email, remainingTickets, userTickets)

		if val {
			fmt.Printf("thank you %v %v for booking the tickets. You will get the confirmation shortly to %v\n", firstName, lastName, email)

			user.UpdateUserDetails(db, firstName, lastName, email, userTickets, cityId)

			remainingTickets = remainingTickets - userTickets
			fmt.Println("abcdabcd")
			cities.UpdateRemainingTickets(db, remainingTickets, cityId)

			flag = true

			waitGroup.Add(1)
			go sendTickets(firstName, email, userTickets)
			continue

		} else {
			flag = false
			continue

		}
	}
	// waitGroup.Wait()

}

func sendTickets(firstName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("Hi %v,\n you have booked %v tickets.\n", firstName, userTickets)
	fmt.Println("###############")
	fmt.Printf("%v The ticket confirmation has been sent to the Email: %v\n ", ticket, email)
	fmt.Println("###############")
	waitGroup.Done()
}
