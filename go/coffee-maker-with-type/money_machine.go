package main

import "fmt"

const (
	CURRENCY = "$"
)

var COIN_VALUES = map[string]float64{
	"quarters": 0.25,
	"dimes":    0.10,
	"nickles":  0.05,
	"pennies":  0.01,
}

var profit float64 = 0

type MoneyMachine struct{}

func (m MoneyMachine) report() {
	fmt.Printf("Money: %s %.2f\n", CURRENCY, profit)
}

func process_coins() float64 {
	var money_received float64 = 0
	for k, v := range COIN_VALUES {
		fmt.Printf("How many %s?: ", k)
		var num float64
		fmt.Scanln(&num)
		money_received += num * v
	}
	return money_received
}

func (m MoneyMachine) make_payment(cost float64) bool {
	money_received := process_coins()
	if money_received >= cost {
		change := money_received - cost
		fmt.Printf("Here is %s%.2f in change.\n", CURRENCY, change)
		profit += cost
		return true
	}
	fmt.Println("Sorry that's not enough money. Money refunded.")
	return false
}
