package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Hello, Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets, and", remainingTickets, "is still available.")
	fmt.Println("Get your tickets here to attend!")

	
}
