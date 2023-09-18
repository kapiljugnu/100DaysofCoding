package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Contains[S ~[]E, E comparable](s S, v E) bool {
	exist := false
	for _, w := range s {
		if w == v {
			exist = true
		}
	}
	return exist

}

func main() {

	word_list := []string{"mouse", "house", "fan", "van"}

	lives := 6

	source := rand.NewSource(time.Now().UnixNano())
	rand_instance := rand.New(source)
	rand_num := rand_instance.Intn(len(word_list))

	word := word_list[rand_num]

	var display []string
	for range word {
		display = append(display, "_")
	}

	end_of_game := false

	for !end_of_game {
		fmt.Println("Guess the letter: ")
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess)

		if Contains(display, guess) {
			fmt.Println("You've already guessed", guess)
		} else {

			for i, w := range word {
				if guess == string(w) {
					display[i] = guess
				}
			}

			t := strings.Split(word, "")
			if !Contains(t, guess) {
				fmt.Println("You guessed", guess, ", that's not in the word. You lose a life")
				lives--
				if lives == 0 {
					end_of_game = true
					fmt.Println(("You lose."))
				}
			}

			if !Contains(display, "_") {
				end_of_game = true
				fmt.Println("You win.")
			}
		}

		fmt.Println(strings.Join(display, ""))
	}
}
