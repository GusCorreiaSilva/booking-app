package main

import (
	"fmt"
	"strings"
)

func main() {
	//2 pontos antes do '=' significa que é uma variavel comum, que não foi definido um valor previo.
	//Ex - var conferenceName string = " " 
  conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get our tickets here to attend")
	bookings := []string{}


	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint


		//ask user for their name
		fmt.Println("enter your first name")
		fmt.Scan(&firstName)

		//ask user for their lastname
		fmt.Println("enter your last name")
		fmt.Scan(&lastName)

		//ask user for their email
		fmt.Println("enter your best email address")
		fmt.Scan(&email)
		
		//ask user how many tickets they want
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets{
			remainingTickets -= userTickets
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets,email)
			fmt.Printf("%v tickets are still available\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var name = strings.Fields(booking)
				firstNames = append(firstNames, name[0])
			}
			fmt.Printf("The first names of bookings are %v\n",firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

			}else{
				fmt.Printf("Sorry we don't have %v tickets, we have only %v tickets remaining\n",userTickets, remainingTickets)
			}

		}
	
}
