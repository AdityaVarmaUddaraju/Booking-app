package main

import "fmt"

func main() {
	// variables in go lang
	var movieName = "After Hours"
	const totalTicketsAvailable = 50
	var remainingTickets = 50

	// Display Welcome message for the users
	// We can send data to standard output using fmt.Print func in go lang
	fmt.Printf("Welcome to %v Movie Ticket Booking App\n",movieName)
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", totalTicketsAvailable, remainingTickets)
	fmt.Println("Please book your desired seats for the film")

	// Using datatypes in go
	var userName string = "Roark"
	var userTickets int = 1

	fmt.Printf("User %v bought %v tickets for the film\n", userName, userTickets)

	// Find the datatype of a variable
	fmt.Printf("Data type of variable userName is %T and userTickets is %T\n", userName, userTickets)

	// Different datatype and their limitations
	var unsignedInteger uint8 = 12
	var signedInteger16 int16 = 25
	var signedInteger8 int8 = -100

	// Go lang throws an error for mismatched type operations
	var unsignedResult  = (signedInteger16) * int16(signedInteger8)

	fmt.Printf("Result is %v\n",unsignedResult)
	fmt.Printf("unsigneg integer type is %T\n", unsignedInteger)

	//syntax sugar
	variable:=12 // this syntax cannot be used to declare constants
	fmt.Printf("type of variable is %T and value is %v\n",variable, variable)

}