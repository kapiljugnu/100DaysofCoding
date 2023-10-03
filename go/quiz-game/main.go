package main

import (
	"fmt"
)

func main() {
	quiz := QuizBrain{}
	for quiz.still_has_question() {
		quiz.next_question()
	}

	fmt.Println("You've completed the quiz.")
	fmt.Printf("Your final score was: %d/%d", quiz.quiz_score(), len(question_data))
}
