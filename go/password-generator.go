package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var letters []string
	for i := 65; i <= 90; i++ {
		letters = append(letters, string(i))
	}
	for i := 97; i <= 122; i++ {
		letters = append(letters, string(i))
	}

	var numbers []string
	for num := 0; num < 10; num++ {
		numbers = append(numbers, strconv.Itoa(num))
	}

	var symbols []string
	for i := 33; i <= 47; i++ {
		symbols = append(symbols, string(i))
	}

	fmt.Println("Welcome to the goPassword Generator!")
	fmt.Println("How many letters would you like in your password?")
	var nr_letters int
	fmt.Scanln(&nr_letters)
	fmt.Println("How many symbols would you like?")
	var nr_symbols int
	fmt.Scanln(&nr_symbols)
	fmt.Println("How many numbers would you like?")
	var nr_numbers int
	fmt.Scanln(&nr_numbers)

	x := rand.NewSource(time.Now().UnixNano())
	rand_instance := rand.New(x)
	var password_list []string
	for i := 0; i < nr_letters; i++ {
		num := rand_instance.Intn(len(letters))
		password_list = append(password_list, letters[num])
	}

	for i := 0; i < nr_symbols; i++ {
		num := rand_instance.Intn(len(symbols))
		password_list = append(password_list, symbols[num])
	}

	for i := 0; i < nr_numbers; i++ {
		num := rand_instance.Intn(len(numbers))
		password_list = append(password_list, numbers[num])
	}

	rand_instance.Shuffle(len(password_list), func(i, j int) {
		password_list[i], password_list[j] = password_list[j], password_list[i]
	})

	fmt.Println(strings.Join(password_list, ""))
}
