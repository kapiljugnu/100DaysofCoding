package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the tip calculator.")
	fmt.Print("What was the total bill? $")
	var total_bill float32
	fmt.Scanln(&total_bill)
	fmt.Print("What percentage tip would you like to give? 10, 12, or 15? ")
	var tip_amount int
	fmt.Scanln(&tip_amount)
	fmt.Print("How many people to split the bill? ")
	var people int
	fmt.Scanln(&people)
	amount := total_bill / float32(people) * (1 + float32(tip_amount)/100)
	fmt.Printf("Each person should pay: $%.2f\n", amount)

}
