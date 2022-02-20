package main

import (
	"fmt"
	"strings"
)

func main() {

	// variables in go lang
	var movieName = "After Hours"
	const totalTicketsAvailable int = 50
	var remainingTickets int = 50

	greetUser(movieName, totalTicketsAvailable, remainingTickets)

	bookings := []string{}

	for {
		firstName, lastName, email, ticketsBooked := getUserInput()

		validName, validEmail, validTicketNumber := validateUser(firstName, lastName, email, ticketsBooked)

		skip := false

		if validName && validEmail && validTicketNumber {

			ticketsAvailable := checkTicketAvailability(remainingTickets, ticketsBooked)

			for !ticketsAvailable{
				fmt.Printf("We have only %v tickets available\n", remainingTickets)
				fmt.Println("Enter Number of tickets required")
				fmt.Scan(&ticketsBooked)
				validTicketNumber = validateTickets(ticketsBooked)
				if !validTicketNumber {
					fmt.Println("You have entered invalid ticket number")
					skip = true
					break
				}
				ticketsAvailable = checkTicketAvailability(remainingTickets, ticketsBooked)
			}

			if skip {
				continue
			}
			
			remainingTickets -= ticketsBooked

			bookings = append(bookings, firstName+" "+lastName)
		
			fmt.Printf("Thank you %v, for booking %v tickets. Your tickets will be sent to %v email\n",firstName, ticketsBooked, email)
			fmt.Printf("Remaining tickets available for %v are %v\n", movieName, remainingTickets)
		
			printBookingNames(bookings)

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

// func definition in go
func greetUser(movieName string, totalTicketsAvailable int, remainingTickets int) {
	// Display Welcome message for the users
	// We can send data to standard output using fmt.Print func in go lang
	fmt.Printf("Welcome to %v Movie Ticket Booking App\n",movieName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", totalTicketsAvailable, remainingTickets)
	fmt.Println("Please book your desired seats for the film")
}

// functions in go can return multiple outputs
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var ticketsBooked int

	fmt.Println("Enter your first name")
	// User input from standard input can be read using fmt.Scan
	fmt.Scan(&firstName) // '&' indicates reference to the variable firstName
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)	
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets required")
	fmt.Scan(&ticketsBooked)

	return firstName, lastName, email, ticketsBooked
}

func validateName(fname string, lname string) bool {
	return len(fname) > 1 && len(lname) > 1
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func validateTickets(tickets int) bool {
	return tickets > 0
}

func validateUser(fname string, lname string, email string, tickets int) (bool, bool, bool) {
	validName := validateName(fname, lname)
	validEmail := validateEmail(email)
	validTicketNumber := validateTickets(tickets)

	return validName, validEmail, validTicketNumber
}

func checkTicketAvailability(remainingTickets int, ticketsBooked int) bool {
	return remainingTickets >= ticketsBooked
}

func printBookingNames(bookings []string) {
	var booking_names []string
		
	for _,name := range bookings {
		full_name := strings.Split(name," ")
		booking_names = append(booking_names, full_name[0])
	}

	fmt.Printf("Current Bookings are %v\n",booking_names)
}