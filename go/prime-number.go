package main

import (
	"fmt"
)

func prime_checker(number int) {
	is_prime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			is_prime = false
		}
	}
	if is_prime {
		fmt.Println("It's a prime number")
	} else {
		fmt.Println("It's not a prime number")
	}
}

func main() {
	fmt.Print("Check this number: ")
	var number int
	fmt.Scanln(&number)
	prime_checker(number)
}
