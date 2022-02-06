package console

import (
	"fmt"

	"stvbyr.tech/go/booking-app/user"
)

func PrintNames(bookings []user.UserData) {
	for _, booking := range bookings {
		first := booking
		fmt.Printf("%v booked\n", first)
	}
}

func PrintGreetings(remainingTickets uint, conferenceName string, conferenceTickets uint) {
	// fmt.Printf("conferenceTickets is %T.\n", conferenceTickets)
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")
}

func PrintSuccessMessage(firstName string, lastName string, userTickets uint, email string, remainingTickets uint) {
	fmt.Printf("Thank you %v %v for booking %v number of tickets. You will receive an email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("we have %v tickets left.\n", remainingTickets)
}

func UserPrompt(firstName *string, lastName *string, email *string, userTickets *uint) {
	// important: you have to give it a pointer to wait for input
	fmt.Println("Enter your first name: ")
	fmt.Scan(firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(email)
	fmt.Println("Enter your amount of tickets: ")
	fmt.Scan(userTickets)
}
