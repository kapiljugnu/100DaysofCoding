package main

import (
	"fmt"
)

func main() {
	fmt.Print("Which number do you want to check? ")
	var number int
	fmt.Scanln(&number)
	if number%2 == 0 {
		fmt.Println("This is an even number.")
	} else {
		fmt.Println("This is an odd number.")
	}
}
