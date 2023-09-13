package main

import (
	"fmt"
)

func main() {
	fmt.Print("Which year do you want to check? ")
	var year int
	fmt.Scanln(&year)
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				fmt.Println("Leap year.")
			} else {
				fmt.Println("Not leap year.")
			}
		} else {
			fmt.Println("Leap year.")
		}
	} else {
		fmt.Println("Not leap year.")
	}
}
