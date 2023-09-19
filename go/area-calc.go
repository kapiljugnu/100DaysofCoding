package main

import (
	"fmt"
  "math"
)

// float64 because we have to round up to largest number
func paint_calc(height, width, cover float64) {
  cans := height * width / cover
	cans = math.Ceil(cans)
	fmt.Printf("You'll need %.f cans of paint.\n", cans)
}

func main() {
	fmt.Print("Height of wall: ")
	var test_h float64
	fmt.Scanln(&test_h)
	fmt.Print("Width of wall: ")
	var test_w float64
	fmt.Scanln(&test_w)
	coverage := float64(5)
	paint_calc(test_h, test_w, coverage)
}