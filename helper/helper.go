package helper

import (
	"fmt"
	"strings"
)

func IsUserInputValid(firstName string, lastName, email string, userTickets uint) bool {
	if len(firstName) < 3 {
		fmt.Println("First name needs at least 2 characters")
		return false
	}
	if len(lastName) < 3 {
		fmt.Println("Last name needs at least 2 characters")
		return false
	}
	if !strings.Contains(email, "@") {
		fmt.Println("Email must contain @")
		return false
	}
	if userTickets < 1 {
		fmt.Println("You can book at least 1 ticket")
		return false
	}

	return true
}

func DoWeHaveTicketsRemaining(remainingTickets uint) bool {
	ticketsRemaining := remainingTickets > 0
	if !ticketsRemaining {
		println("No tickets left.")
		return false
	}

	return true
}

func UserWantsMoreTicketsThanRemaining(userTickets uint, remainingTickets uint) bool {
	userWantsMoreTicketsThanRemaining := userTickets > remainingTickets
	if userWantsMoreTicketsThanRemaining {
		fmt.Printf("You want to book %v tickets but there are inly %v tickets left.\n", userTickets, remainingTickets)
		fmt.Println("Please try again.")
		return true
	}

	return false
}
