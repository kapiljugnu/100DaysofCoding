package main

import (
	"fmt"
)

func main() {
	var total int
	for i := 1; i < 101; i++ {
		if total % 2 == 0 {
			total += i
		}
	}
	fmt.Println(total)
}