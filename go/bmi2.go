package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("enter your height in m: ")
	var height float64
	fmt.Scanln(&height)
	fmt.Print("enter your weight in kg: ")
	var weight float64
	fmt.Scanln(&weight)
	bmi := weight / math.Pow(height, 2)
	rounded_bmi := int(bmi)
	if bmi < 18.5 {
		fmt.Println("Your BMI is", rounded_bmi, ", you are underweight.")
	} else if bmi < 25 {
		fmt.Println("Your BMI is", rounded_bmi, ", you have a normal weight.")
	} else if bmi < 30 {
		fmt.Println("Your BMI is", rounded_bmi, ", you are slightly overweight.")
	} else if bmi < 35 {
		fmt.Println("Your BMI is", rounded_bmi, ", you are obese.")
	} else {
		fmt.Println("Your BMI is", rounded_bmi, ", you are clinically obese.")
	}
}
