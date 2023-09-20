package main

import (
	"fmt"
	"strings"
)

type dType float64

func add(n1, n2 dType) dType {
	return n1 + n2
}

func subtract(n1, n2 dType) dType {
	return n1 - n2
}

func multiply(n1, n2 dType) dType {
	return n1 * n2
}

func divide(n1, n2 dType) dType {
	return n1 / n2
}

func main() {
	operations := map[string]func(dType, dType) dType{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	fmt.Print("Enter the first number?: ")
	var num1 dType
	fmt.Scanln(&num1)

	ops := []string{}
	for key := range operations {
		ops = append(ops, key)
	}
	fmt.Println(strings.Join(ops, ","))
	fmt.Print("Pick an operation from the line above: ")
	var operation_symbol string
	fmt.Scanln(&operation_symbol)

	fmt.Print("Enter the second number?: ")
	var num2 dType
	fmt.Scanln(&num2)

	fn := operations[operation_symbol]
	answer := fn(num1, num2)

	fmt.Printf("%.f %s %.f = %.f\n", num1, operation_symbol, num2, answer)
}
