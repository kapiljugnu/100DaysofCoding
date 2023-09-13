package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Python Pizza Deliveries!")
	fmt.Print("What size pizza do you want? S, M, or L ")
	var size string
	fmt.Scanln(&size)
	fmt.Print("Do you want pepperoni? Y or N ")
	var add_pepperoni string
	fmt.Scanln(&add_pepperoni)
	fmt.Print("Do you want extra cheese? Y or N ")
	var extra_cheese string
	fmt.Scanln(&extra_cheese)

	bill := 0
	if size == "S" {
		bill += 15
	} else if size == "M" {
		bill += 20
	} else {
		bill += 25
	}

	if add_pepperoni == "Y" {
		if size == "S" {
			bill += 2
		} else {
			bill += 3

		}
	}

	if extra_cheese == "Y" {
		bill += 1
	}

	fmt.Println("Your final bill is: $", bill)
}
