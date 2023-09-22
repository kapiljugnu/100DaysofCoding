package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_SCORE  = 21
	BLACK_JACK = 999999
	ACE        = 11
)

func deal_card() int {
	cards := []int{ACE, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	source := rand.NewSource(time.Now().UnixNano())
	rand_instance := rand.New(source)

	card := rand_instance.Intn(len(cards))
	return card
}

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

func calculate_score(cards []int) ([]int, int) {
	total := sum(cards)

	if len(cards) == 2 && total == MAX_SCORE {
		return cards, BLACK_JACK
	}

	if total > MAX_SCORE && Contains(cards, ACE) {
		cards_with_one := []int{}
		match_found := false
		for _, card := range cards {
			num := card
			if card == ACE && !match_found {
				num = 1
				match_found = true
			}
			cards_with_one = append(cards_with_one, num)
		}
		return cards_with_one, sum(cards_with_one)
	}

	return cards, total
}

func compare(player_score, computer_score int) string {
	if player_score > MAX_SCORE && computer_score > MAX_SCORE {
		return "You went over. You lose."
	}

	if player_score == computer_score {
		return "Draw 🙃"
	} else if computer_score == BLACK_JACK {
		return "Lose, opponent has Blackjack 😱"
	} else if player_score == BLACK_JACK {
		return "Win with a Blackjack 😎"
	} else if player_score > MAX_SCORE {
		return "You went over you lose 😭"
	} else if computer_score > MAX_SCORE {
		return "Opponent went over. You win 😀"
	} else if player_score > computer_score {
		return "You win 😀"
	} else {
		return "You lose 😭"
	}
}

func play_game() {

	fmt.Println("Welcome to Blackjack game.")
	var (
		player_cards,
		computer_cards []int
		computer_score,
		player_score int
	)

	for i := 0; i < 2; i++ {
		player_cards = append(player_cards, deal_card())
		computer_cards = append(computer_cards, deal_card())
	}

	is_game_over := false
	for !is_game_over {
		player_cards, player_score = calculate_score(player_cards)
		computer_cards, computer_score = calculate_score(computer_cards)

		fmt.Printf("Your cards : %v, current score: %d\n", player_cards, player_score)
		fmt.Println("Computer's first card:", computer_cards[0])

		if player_score == BLACK_JACK || computer_score == BLACK_JACK || player_score > MAX_SCORE {
			is_game_over = true
		} else {
			fmt.Print("Type 'y' to get another card, type 'n' to pass: ")
			var another_card string
			fmt.Scanln(&another_card)
			if another_card == "y" {
				player_cards = append(player_cards, deal_card())
			} else {
				is_game_over = true
			}
		}
	}

	for computer_score != BLACK_JACK && computer_score < 17 {
		computer_cards = append(computer_cards, deal_card())
		computer_cards, computer_score = calculate_score(computer_cards)
	}
	fmt.Printf("Your final hand: %v, final score: %d\n", player_cards, player_score)
	fmt.Printf("Computer's final hand: %v, final score: %d\n", computer_cards, computer_score)
	fmt.Println(compare(player_score, computer_score))
}

func main() {
	play_game()
}
