package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to brand name generator.")
  fmt.Println("What is the name of your city?")
  var name, pet string
  fmt.Scanln(&name)
  fmt.Println("What is the name of your pet?")
  fmt.Scanln(&pet)
  fmt.Println("You brand name could be", name, pet)
  
}
