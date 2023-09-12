package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter your height. ")
	var height float64
	fmt.Scanln(&height)
	fmt.Print("Enter your weight. ")
	var weight float64
	fmt.Scanln(&weight)
	bmi := weight / math.Pow(height, 2)
	fmt.Printf("%.2f\n", bmi)
}
