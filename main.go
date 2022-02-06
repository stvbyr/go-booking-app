package main

import (
	"fmt"
	"sync"
	"time"

	"stvbyr.tech/go/booking-app/console"
	"stvbyr.tech/go/booking-app/helper"
	"stvbyr.tech/go/booking-app/user"
)

const conferenceTickets = 50
const conferenceName = "Go Conference"

var remainingTickets uint = 50

// arrays
// var array = [50]string{"Nana", "Nicole", "Peter"}
// var bookings [50]string
// slices
// var bookings []string
// var bookings = []string{}
// maps
// var bookings = make([]map[string]string, 0)
// types / structs
var bookings = []user.UserData{}
var wg = sync.WaitGroup{}

func main() {
	console.PrintGreetings(remainingTickets, conferenceName, conferenceTickets)

	ticketsRemaining := helper.DoWeHaveTicketsRemaining(remainingTickets)

	if !ticketsRemaining {
		return
	}

	var lastName string
	var firstName string
	var email string
	var userTickets uint

	console.UserPrompt(&firstName, &lastName, &email, &userTickets)
	if !helper.IsUserInputValid(firstName, lastName, email, userTickets) {
		return
	}

	if helper.UserWantsMoreTicketsThanRemaining(userTickets, remainingTickets) {
		return
	}

	bookTicket(userTickets, firstName, lastName, email)

	//goroutines / concurrency with the go keyword
	// if the main thread ends the routine terminates as well no matter if it was executed
	// with wait groups we can tell go to wait for sub threads
	// if we had multiple routines the number in Add would have to be increased
	wg.Add(1)
	go sendTickets()

	console.PrintNames(bookings)

	// because this is the end of the main thread we tell it to wait for the sub threads to finish
	wg.Wait()
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// add bookings (array)
	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("Array: %v\n", bookings)
	// fmt.Printf("Length: %v\n", len(bookings))

	// add bookings (slices)
	// bookings = append(bookings, firstName+" "+lastName)

	// add bookings (maps)
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// add bookings (types / structs)
	userData := user.UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	console.PrintSuccessMessage(firstName, lastName, userTickets, email, remainingTickets)
}

func sendTickets() {
	time.Sleep(10 * time.Second)
	fmt.Println("Sending ticket.")
	// we can signal that a sub thread has finished
	wg.Done()
}
