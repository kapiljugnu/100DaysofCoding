package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Treasure Island.")
	fmt.Println("Your mission is to find the treasure.")
	fmt.Println("You're at a cross road. Where do you want to go? Type \"left\" or \"right\"")
	var choice1 string
	fmt.Scanln(&choice1)
	choice1 = strings.ToLower(choice1)
	if choice1 == "left" {
		fmt.Println("You've come to a lake. There is an island in the middle of the lake. Type \"wait\" to wait for a boat. Type \"swim\" to swim across.")
		var choice2 string
		fmt.Scanln(&choice2)
		choice2 = strings.ToLower(choice2)
		if choice2 == "wait" {
			fmt.Println("You arrive at the island unharmed. There is a house with 3 doors. One red, one yellow, and one blue. Which colour do you choose?")
			var choice3 string
			fmt.Scanln(&choice3)
			choice3 = strings.ToLower(choice3)
			if choice3 == "red" {
				fmt.Println("It's a room full of fire. Game Over.")
			} else if choice3 == "yellow" {
				fmt.Println("You found the treasure! You win!")
			} else if choice3 == "blue" {
				fmt.Println("You enter a room of beasts. Game Over.")
			} else {
				fmt.Println("You chose a door that doesn't exist. Game Over.")
			}
		} else {
			fmt.Println("You get attcked by an angry trout. Game Over.")
		}
	} else {
		fmt.Println("You fell into a hole. Game Over.")
	}
}
