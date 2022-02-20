package main

import "fmt"

func main() {
	// variables in go lang
	var movieName = "After Hours"
	const totalTicketsAvailable uint = 50
	var remainingTickets uint = 50

	// static length arrays
	var fixedLengthArray [10]string 
	fixedLengthArray[2] = "element3" // elements can be assigned using index
	fmt.Printf("Whole array is %v\n",fixedLengthArray)
	fmt.Printf("Datatype of array is %T\n", fixedLengthArray)
	fmt.Printf("first element of array %v\n", fixedLengthArray[0])
	fmt.Printf("length of array is %v\n",len(fixedLengthArray))

	// dynamic length arrays
	// variable length arrays are called slices in go
	slices := []string{}
	slices = append(slices, "element1") // elements can be assigned using append()
	fmt.Printf("Whole slice is %v\n",slices)
	fmt.Printf("Datatype of slice is %T\n", slices)
	fmt.Printf("first element of slice is %v\n", slices[0])
	fmt.Printf("length of array slice %v\n",len(slices))


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
	fmt.Printf("Current Bookings are %v\n",bookings)
}