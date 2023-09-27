package main

import (
	"fmt"
)

type Ingredients struct {
	water,
	coffee,
	milk int
}
type CoffeeType struct {
	ingredients Ingredients
	cost        float64
}

var MENU = map[string]CoffeeType{
	"espresso": {
		ingredients: Ingredients{
			water:  50,
			coffee: 18,
			milk:   0,
		},
		cost: 1.5,
	},

	"latte": {
		ingredients: Ingredients{
			water:  200,
			milk:   150,
			coffee: 24,
		},
		cost: 2.5,
	},

	"cappuccino": {
		ingredients: Ingredients{
			water:  250,
			milk:   100,
			coffee: 24,
		},
		cost: 3.0,
	},
}

var resources = Ingredients{
	water:  300,
	milk:   200,
	coffee: 100,
}

func print_report(sale_amount float64) {
	fmt.Printf("Water: %dml\n", resources.water)
	fmt.Printf("Milk: %dml\n", resources.milk)
	fmt.Printf("Coffee: %dg\n", resources.coffee)
	fmt.Println("Money: $", sale_amount)
}

func check_resources(coffee string) string {
	coffee_type := MENU[coffee]
	ingredients := coffee_type.ingredients
	if resources.water < ingredients.water {
		return "water"
	} else if resources.milk < ingredients.milk {
		return "milk"
	} else if resources.coffee < ingredients.coffee {
		return "coffee"
	}

	return ""
}

func process_coins() float64 {
	fmt.Print("how many quaters?: ")
	var quaters float64
	fmt.Scanln(&quaters)
	quaters = 0.25 * quaters
	fmt.Print("how many dimes?: ")
	var dimes float64
	fmt.Scanln(&dimes)
	dimes = 0.10 * dimes
	fmt.Print("how many nickles?: ")
	var nickles float64
	fmt.Scanln(&nickles)
	nickles = 0.05 * nickles
	fmt.Print("how many pennies?: ")
	var pennies float64
	fmt.Scanln(&pennies)
	pennies = 0.01 * pennies

	return quaters + dimes + nickles + pennies
}

func reduce_resources(coffee string) {
	coffee_type := MENU[coffee]
	ingredients := coffee_type.ingredients
	if resources.water > 0 {
		resources.water -= ingredients.water
	}

	if resources.milk > 0 {
		resources.milk -= ingredients.milk
	}

	if resources.coffee > 0 {
		resources.coffee -= ingredients.coffee
	}
}

func main() {
	continue_serving := true
	var total_sale float64
	for continue_serving {
		fmt.Print("What would you like? (espresso/latte/cappuccino): ")
		var coffee_type string
		fmt.Scanln(&coffee_type)

		if coffee_type == "off" {
			continue_serving = false
		} else if coffee_type == "report" {
			print_report(total_sale)
		} else {
			insufficient_resource := check_resources(coffee_type)

			if insufficient_resource != "" {
				fmt.Printf("Sorry there is not enough %s.\n", insufficient_resource)
			} else {
				fmt.Println("Please insert coins.")
				total := process_coins()

				cost := MENU[coffee_type].cost
				if cost > total {
					fmt.Println("Sorry that's not enough money. Money refunded.")
				} else {
					change := total - cost
					total_sale += cost
					reduce_resources(coffee_type)
					fmt.Printf("Here is $%.2f in change\n", change)
					fmt.Printf("Here is your %s ☕️. Enjoy!\n", coffee_type)
				}
			}
		}
	}
}
