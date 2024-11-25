package user

import (
	"conferenceTicketBooking/types"
	"database/sql"
	"fmt"
	"log"
)

func UpdateUserDetails(db *sql.DB, firstName string, lastName string, email string, tickets uint, cityId int) {
	var usersList = []types.User{}
	query := "INSERT INTO user(firstName, lastName, email, tickets, cityId) VALUES (?, ?, ?, ?, ?)"

	_, err := db.Exec(query, firstName, lastName, email, tickets, cityId)
	if err != nil {
		log.Fatal("Query was not able be executed")
	}
	println("You are now added to our Guest list")
	//////////////////////////////////////////////////////////////////////
	query = "SELECT cityName FROM city where cityId=?"
	res, err := db.Query(query, cityId)

	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	var ct string
	res.Scan(&ct)
	fmt.Printf("%s Conference\n", ct)

	rows, err := db.Query("SELECT Id, firstName, lastName, email, tickets FROM user where user.cityId=?", cityId)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Tickets)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		usersList = append(usersList, user)

	}
	for _, user := range usersList {
		fmt.Printf("%v. %v %v\n", user.Id, user.FirstName, user.Tickets)
	}

}
