package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := rand.NewSource(time.Now().UnixNano())
	y := rand.New(x)
	num := y.Intn(2)
	if num == 1 {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}
}
