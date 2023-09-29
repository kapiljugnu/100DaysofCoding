package main

import (
	"fmt"
)

func main() {
	continue_serving := true
	menu := Menu{}
	coffee_maker := CoffeeMaker{}
	money_machine := MoneyMachine{}
	for continue_serving {
		fmt.Printf("What would you like? (%s): ", menu.getItems())
		var choice string
		fmt.Scanln(&choice)

		if choice == "off" {
			continue_serving = false
			return
		}

		if choice == "report" {
			coffee_maker.report()
			money_machine.report()
			continue
		}

		drink, err := menu.find_drink(choice)
		if err != nil {
			panic(err)
		}
		if coffee_maker.is_resource_sufficient(drink) && money_machine.make_payment(drink.cost) {
			coffee_maker.make_coffee(drink)
		}
	}
}
