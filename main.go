package main

import (
	"fmt"
	
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get our tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	fmt.Println("enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter your best email address")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	
	remainingTickets -= userTickets

	fmt.Printf("thank you  %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets,email)
	fmt.Printf("%v tickets are still available\n", remainingTickets)
}
