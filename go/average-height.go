package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input a list of student heights ")
	reader := bufio.NewReader(os.Stdin)
	student_heights, _ := reader.ReadString('\n')
	student_heights = strings.TrimSpace(student_heights)
	heights := strings.Split(student_heights, " ")

	var height_total int
	var count int
	for _, h := range heights {
		num, err := strconv.Atoi(h)
		if err != nil {
			panic(err)
		}
		height_total += num
		count++
	}

	fmt.Println(height_total, count)
	avg := height_total / count
	fmt.Println(avg)
}
