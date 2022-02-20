package main

import (
	"fmt"
	"strings"
)

func main() {

	// for loop
	for i:=0; i<10; i++ {
		fmt.Printf("number is %v\n", i)
	}

	// for each
	slice := []string{"element1","element2"}
	// syntax index, element := range list {}
	for _,element := range slice { 
		fmt.Printf("element is %v\n", element)
	}

	// variables in go lang
	var movieName = "After Hours"
	const totalTicketsAvailable uint = 50
	var remainingTickets uint = 50

	// Display Welcome message for the users
	// We can send data to standard output using fmt.Print func in go lang
	fmt.Printf("Welcome to %v Movie Ticket Booking App\n",movieName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", totalTicketsAvailable, remainingTickets)
	fmt.Println("Please book your desired seats for the film")

	var firstName string
	var lastName string
	var email string
	var ticketsBooked uint

	bookings := []string{}

	for {
		fmt.Println("Enter your first name")
		// User input from standard input can be read using fmt.Scan
		fmt.Scan(&firstName) // '&' indicates reference to the variable firstName
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)	
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets required")
		fmt.Scan(&ticketsBooked)
	
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
	}

	
}