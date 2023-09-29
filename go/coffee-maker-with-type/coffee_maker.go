package main

import "fmt"

var resources Ingredients

func init() {
	resources = Ingredients{
		water:  300,
		milk:   200,
		coffee: 100,
	}
}

type CoffeeMaker struct{}

func (CoffeeMaker) report() {
	fmt.Printf("Water: %dml\n", resources.water)
	fmt.Printf("Milk: %dml\n", resources.milk)
	fmt.Printf("Coffee: %dg\n", resources.coffee)
}

func (CoffeeMaker) is_resource_sufficient(drink MenuItem) bool {
	ingredients := drink.ingredients
	if ingredients.water > resources.water {
		fmt.Println("Sorry there is not enough water.")
		return false
	} else if ingredients.milk > resources.milk {
		fmt.Println("Sorry there is not enough milk.")
		return false
	} else if ingredients.coffee > resources.coffee {
		fmt.Println("Sorry there is not enough coffee.")
		return false
	}
	return true
}

func (CoffeeMaker) make_coffee(order MenuItem) {
	resources.water -= order.ingredients.water
	resources.coffee -= order.ingredients.coffee
	resources.milk -= order.ingredients.milk
	fmt.Printf("Here is your %s ☕️. Enjoy!\n", order.name)
}
