package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [] string 

	fmt.Printf("Hello, Welcome to %v booking application! \n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v is still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for {
		var firstName string 
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you would like:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("There are %v remaining tickets for %v \n", remainingTickets, conferenceName)

		firstNames := [] string {}

		// to iterate through a slice/array use range as it provides index and value for each el
		for _, el := range bookings {
			var names = strings.Fields(el)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("These are all the bookings: %v \n", bookings)
	}

	
}

