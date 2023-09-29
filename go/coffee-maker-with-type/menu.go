package main

import "errors"

type Ingredients struct {
	water  int
	milk   int
	coffee int
}

type MenuItem struct {
	name        string
	cost        float64
	ingredients Ingredients
}

var menus []MenuItem

func init() {
	menus = []MenuItem{
		{
			name: "espresso",
			cost: 1.5,
			ingredients: Ingredients{
				water:  50,
				coffee: 18,
				milk:   0,
			},
		},
		{
			name: "latte",
			cost: 2.5,
			ingredients: Ingredients{
				water:  200,
				milk:   150,
				coffee: 24,
			},
		},
		{
			name: "cappuccino",
			cost: 3.0,
			ingredients: Ingredients{
				water:  250,
				milk:   100,
				coffee: 24,
			},
		},
	}
}

type Menu struct {}

func (Menu) getItems() string {
	options := ""
	for _, item := range menus {
		options += (item.name + "/")
	}
	return options
}

var ErrDrinkNotAvailable = errors.New("Sorry that item is not available.")

func (Menu) find_drink(order_name string) (MenuItem, error) {
	for _, item := range menus {
		if item.name == order_name {
			return item, nil
		}
	}
	return MenuItem{}, ErrDrinkNotAvailable
}
