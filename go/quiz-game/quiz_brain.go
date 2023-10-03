package main

import (
	"fmt"
	"strings"
)

var (
	score           = 0
	question_number = 0
)

type QuizBrain struct{}

func (QuizBrain) still_has_question() bool {
	return question_number < len(question_data)
}

func (QuizBrain) next_question() {
	current_question := question_data[question_number]
	question_number++
	fmt.Printf("Q.%d: %s (True/False): ", question_number, current_question.text)
	var answer string
	fmt.Scanln(&answer)
	check_answer(answer, current_question.answer)
}

func (QuizBrain) quiz_score() int {
	return score
}

func check_answer(user_answer, correct_answer string) {
	if strings.ToLower(user_answer) == strings.ToLower(correct_answer) {
		fmt.Println("You got it right!")
		score++
	} else {
		fmt.Println("That's wrong.")
	}

	fmt.Printf("The correct answer was: %s\n", correct_answer)
	fmt.Printf("The current score is %d/%d\n", score, question_number)
	fmt.Println("")
}
