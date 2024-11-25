package cities

import (
	"conferenceTicketBooking/types"
	"database/sql"
	"fmt"
	"log"
)

// var remainingTickets uint
// var id int = 0
// var cities = []string{"Berlin", "Boston", "Chicago", "New York"}
// var citiesMap = map[int]uint{
// 	1: 50,
// 	2: 80,
// 	3: 90,
// 	4: 40,
// }

// func GetCities() (uint, int) {

// 	var userInputcity int
// 	for {
// 		fmt.Println("Please enter the index of one of the following cities ")
// 		for idx, city := range cities {
// 			fmt.Printf("%d. %s\n", idx+1, city)
// 		}

// 		fmt.Print("> ")
// 		fmt.Scan(&userInputcity)

// 		remainingTickets = citiesMap[userInputcity]
// 		fmt.Println(userInputcity, remainingTickets)

// 		if remainingTickets > 0 {
// 			return remainingTickets, userInputcity
// 		} else {

// 			fmt.Println("Sorry, the tickects are sold out. kindy try another city?")
// 			continue

// 		}

// 	}

// }
func UpdateRemainingTickets(db *sql.DB, remainingTickets uint, Id int) {
	query := "UPDATE city SET capacity = ? WHERE cityId = ?"

	_, err := db.Exec(query, remainingTickets, Id)
	if err != nil {
		log.Fatal("error updating city capacity: %v", err)
	}

}

// func GetRemainingTiclets()

func GetCities(db *sql.DB) (uint, int) {
	var selectedCity types.City
	var userInputcity int
	var Id int
	// Query to get all rows from the city table
	rows, err := db.Query("SELECT cityId, cityName, capacity FROM city")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	var cityList []types.City

	// Iterate over the rows
	for rows.Next() {
		var ct types.City
		// Scan values from the row into the city struct
		err := rows.Scan(&ct.ID, &ct.Name, &ct.Capacity)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}

		cityList = append(cityList, ct)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}

	for {
		fmt.Println("Please enter the index of one of the following cities ")

		for _, city := range cityList {
			fmt.Printf("%v. %v\n", city.ID, city.Name)
		}
		fmt.Print("> ")
		fmt.Scan(&userInputcity)

		for _, ct := range cityList {
			if ct.ID == userInputcity {
				selectedCity = ct
				break
			}
		}

		if selectedCity.Capacity > 0 {
			fmt.Printf("There are %v tickets in %v\n", selectedCity.Capacity, selectedCity.Name)
			break
		} else {

			fmt.Println("Sorry, the tickects are sold out. kindly try another city?")
			continue

		}

	}
	Id = selectedCity.ID
	return selectedCity.Capacity, Id

}
