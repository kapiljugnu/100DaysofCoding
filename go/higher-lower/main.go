package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(numbers []int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	exist := false
	for _, w := range s {
		if w == v {
			exist = true
		}
	}
	return exist

}

func pick_random_account() Account {
	source := rand.NewSource(time.Now().UnixNano())
	rand_instance := rand.New(source)

	idx := rand_instance.Intn(len(Data))
	return Data[idx]
}

func check_answer(guess string, a_followers, b_followers int) bool {
	if a_followers > b_followers {
		return guess == "A"
	} else {
		return guess == "B"
	}
}

func play_game() {
	var score int
	celebrity_a := pick_random_account()

	var is_game_over bool
	for !is_game_over {
		celebrity_b := pick_random_account()
		fmt.Println("****************************")
		fmt.Printf("Compare A: %s, a %s, from %s\n", celebrity_a.name, celebrity_a.description, celebrity_a.country)
		fmt.Printf("Against B: %s, a %s, from %s\n", celebrity_b.name, celebrity_b.description, celebrity_b.country)

		fmt.Println("Who has more followers? Type 'A' or 'B': ")
		var guess string
		fmt.Scanln(&guess)

		a_follower := celebrity_a.follower_count
		b_follower := celebrity_b.follower_count
		is_correct := check_answer(guess, a_follower, b_follower)
		fmt.Println("****************************")
		if is_correct {
			score += 1
			celebrity_a = celebrity_b
			fmt.Println("You're right! Current score:", score)
		} else {
			is_game_over = true
			fmt.Println("Sorry, that' wrong. Final score:", score)
		}
	}
}

func main() {
	play_game()
}
