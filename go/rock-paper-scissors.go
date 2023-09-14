package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rock := "🪨"
	paper := "📄"
	scissors := "✂️"

	tools := []string{rock, paper, scissors}

	fmt.Println("What do you choose? Type 0 for Rock, 1 for Paper or 2 for Scissors.")
	var user_choice int
	fmt.Scanln(&user_choice)

	if user_choice >= 3 || user_choice < 0 {
		fmt.Println("You typed an invalid number, you lose!")
	} else {
		fmt.Println(tools[user_choice])

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		computer_choice := r.Intn(3)

		fmt.Println(tools[computer_choice])

		if user_choice == 0 && computer_choice == 2 {
			fmt.Println("You win!")
		} else if computer_choice == 0 && user_choice == 2 {
			fmt.Println("You lose")
		} else if computer_choice > user_choice {
			fmt.Println("You lose")
		} else if user_choice > computer_choice {
			fmt.Println("You win!")
		} else if computer_choice == user_choice {
			fmt.Println("It's a draw")
		}

	}
}
