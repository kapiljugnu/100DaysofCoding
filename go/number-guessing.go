package main

import (
	"fmt"
	"math/rand"
	"time"
)

var lives int = 5

func reduce_live() {
	lives--
	if lives == 0 {
		fmt.Println("You've run out of guesses, you lose")
	} else {
		fmt.Println("Guess again.")
	}
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100")

	fmt.Print("Choose a difficulty. Type 'easy' or 'hard': ")
	var difficulty string
	fmt.Scanln(&difficulty)

	if difficulty == "easy" {
		lives = 10
	}

	source := rand.NewSource(time.Now().UnixNano())
	rand_instance := rand.New(source)

	computer_number := rand_instance.Intn(100-2) + 1

	for lives > 0 {
		fmt.Printf("You have %d attempts remaining to guess the number.\n", lives)
		var guess int
		fmt.Print("Make a guess: ")
		fmt.Scanln(&guess)
		if guess == computer_number {
			fmt.Println("You got it! The answer was", computer_number)
			lives = 0
		} else if guess > computer_number {
			fmt.Println("Too high.")
			reduce_live()
		} else {
			fmt.Println("Too low.")
			reduce_live()
		}
	}
}
