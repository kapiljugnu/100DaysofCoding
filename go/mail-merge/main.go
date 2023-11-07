package main

import (
	"os"
	"strings"
)

func main() {
	letter_template, err := os.ReadFile("./input/letter.txt")
	if err != nil {
		panic(err)
	}

	invites, err := os.ReadFile("./input/names.txt")
	if err != nil {
		panic(err)
	}

	names := strings.Split(string(invites), "\n")

	for _, name := range names {
		letter := strings.Replace(string(letter_template), "[name]", name, 1)
		err := os.WriteFile("./output/letter_to_"+name+".txt", []byte(letter), 0644)

		if err != nil {
			panic(err)
		}
	}
}
