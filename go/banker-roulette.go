package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give me everybody's names, separated by a comma. ")
	names_string, _ := reader.ReadString('\n')
	names_string = strings.TrimSpace(names_string)
	names := strings.Split(names_string, ",")

	x := rand.NewSource(time.Now().UnixNano())
	y := rand.New(x)
	num := y.Intn(len(names))

	fmt.Println(names[num], "is going to buy the meal today!")
}
