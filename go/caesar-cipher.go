package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var alphabets []string

func init() {
	for i := 97; i < 123; i++ {
		alphabets = append(alphabets, string(i))
	}

	alphabets = append(alphabets, alphabets...)
}

func Index[S ~[]E, E comparable](s S, v E) int {
	for i, w := range s {
		if w == v {
			return i
		}
	}
	return -1

}

func ceaser(start_text string, shift_amount int, cipher_direction string) {
	var end_text string
	direction := "encoded"
	if cipher_direction == "d" {
		direction = "decoded"
		shift_amount *= -1
	}
	for _, char := range start_text {
		letter := string(char)
		position := Index(alphabets, letter)
		if position > -1 {
			new_position := position + shift_amount
			end_text += alphabets[new_position]
		} else {
			end_text += letter
		}
	}

	fmt.Println("Here's the", direction, "result:", end_text)
	fmt.Println("")
}

func main() {
	should_continue := true
	for should_continue {
		fmt.Println("Type 'e' to encrypt, type 'd' to decrypt:")
		var direction string
		fmt.Scanln(&direction)

		fmt.Println("Type your message:")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ToLower(text)

		fmt.Println("Type the shift number:")
		var shift int
		fmt.Scanln(&shift)

		shift = shift % 26

		ceaser(text, shift, direction)

		var restart string
		fmt.Println("Type 'y' if you want to go again. Otherwise type 'n'.")
		fmt.Scanln(&restart)

		if restart == "n" {
			should_continue = false
			fmt.Println("Goodbye")
		}
	}
}
