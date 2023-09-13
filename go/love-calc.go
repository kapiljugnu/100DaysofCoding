package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Love Calculator!")
	fmt.Println("What is your name?")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSpace(name1)
	fmt.Println("What is their name?")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSpace(name2)

	name := strings.ToLower(name1 + name2)

	t := strings.Count(name, "t")
	r := strings.Count(name, "r")
	u := strings.Count(name, "u")
	e := strings.Count(name, "e")

	trueScore := t + r + u + e

	l := strings.Count(name, "l")
	o := strings.Count(name, "o")
	v := strings.Count(name, "v")
	e = strings.Count(name, "e")

	loveScore := l + o + v + e

	score := strconv.Itoa(trueScore) + strconv.Itoa(loveScore)

	int_score, err := strconv.Atoi(score)
	if err != nil {
		panic(err)
	}

	if int_score < 10 || int_score > 90 {
		fmt.Printf("Your score is %d, you go together like coke and mentos.\n", int_score)
	} else if int_score >= 40 && int_score <= 50 {
		fmt.Printf("Your score is %d, you are alright together.\n", int_score)
	} else {
		fmt.Printf("Your score is %d.\n", int_score)
	}

}
