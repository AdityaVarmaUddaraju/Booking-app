package main

import (
	"fmt"
	"strings"
)

func main() {

	// switch statements in go
	theaterName := "After Hours"

	switch theaterName {
		case "After Hours": {
			fmt.Println("Movie name is After Hours")
		}
		case "Avengers", "Spider-Man": {
			fmt.Println("Movie name is Avengers or Spider-Man")
		}
		default: {
			fmt.Println("Unknown movie name")
		}

	}

	// variables in go lang
	var movieName = "After Hours"
	const totalTicketsAvailable int = 50
	var remainingTickets int = 50

	// Display Welcome message for the users
	// We can send data to standard output using fmt.Print func in go lang
	fmt.Printf("Welcome to %v Movie Ticket Booking App\n",movieName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", totalTicketsAvailable, remainingTickets)
	fmt.Println("Please book your desired seats for the film")

	var firstName string
	var lastName string
	var email string
	var ticketsBooked int

	bookings := []string{}

	for {
		fmt.Println("Enter your first name")
		// User input from standard input can be read using fmt.Scan
		fmt.Scan(&firstName) // '&' indicates reference to the variable firstName
		fmt.Printf("first name is %v\n", firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)	
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets required")
		fmt.Scan(&ticketsBooked)

		validName := len(firstName) > 1 && len(lastName) > 1
		validEmail := strings.Contains(email, "@")
		validTicketNumber := ticketsBooked > 0
		skip := false

		if validName && validEmail && validTicketNumber {

			ticketsAvailable := remainingTickets >= ticketsBooked

			for !ticketsAvailable{
				fmt.Printf("We have only %v tickets available\n", remainingTickets)
				fmt.Println("Enter Number of tickets required")
				fmt.Scan(&ticketsBooked)
				validTicketNumber = ticketsBooked > 0
				if !validTicketNumber {
					fmt.Println("You have entered invalid ticket number")
					skip = true
					break
				}
				ticketsAvailable = remainingTickets >= ticketsBooked
			}

			if skip {
				continue
			}
			
			remainingTickets -= ticketsBooked

			bookings = append(bookings, firstName+" "+lastName)
		
			fmt.Printf("Thank you %v, for booking %v tickets. Your tickets will be sent to %v email\n",firstName, ticketsBooked, email)
			fmt.Printf("Remaining tickets available for %v are %v\n", movieName, remainingTickets)
		
			var booking_names []string
		
			for _,name := range bookings {
				full_name := strings.Split(name," ")
				booking_names = append(booking_names, full_name[0])
			}
		
			fmt.Printf("Current Bookings are %v\n",booking_names)

			if remainingTickets == 0 {
				fmt.Println("We are sold out")
				break
			}

		} else {
			if !validName {
				fmt.Println("You have entered invalid name")
			}
			if !validEmail {
				fmt.Println("You have entered invalid email")
			}
			if !validTicketNumber {
				fmt.Println("You have entered invalid ticket number")
			}
		}
	

	}

	
}