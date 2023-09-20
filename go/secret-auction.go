package main

import (
	"fmt"
)

func main() {
	auction := make(map[string]int)
	new_bid := true
	for new_bid {
		fmt.Print("What is your name?: ")
		var name string
		fmt.Scanln(&name)
		fmt.Print("What's your bid?: $")
		var bid_amount int
		fmt.Scanln(&bid_amount)
		auction[name] = bid_amount

		fmt.Print("Are there any other bidders? Type 'y' or 'n' ")
		var restart string
		fmt.Scanln(&restart)

		if restart == "n" {
			new_bid = false
		}
	}

	var (
		max_amount  int
		bidder_name string
	)

	for key, value := range auction {
		if value > max_amount {
			max_amount = value
			bidder_name = key
		}
	}

	fmt.Printf("The winner is %s with a bid of $%d\n", bidder_name, max_amount)
}
