package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input a list of student scores ")
	reader := bufio.NewReader(os.Stdin)
	student_scores, _ := reader.ReadString('\n')
	student_scores = strings.TrimSpace(student_scores)
	scores := strings.Split(student_scores, " ")

  var max int
  for _, score := range scores {
    s, _ := strconv.Atoi(score)
    if s > max {
      max = s
    }
  }
  fmt.Println("The highest score in the class is:", max)
}
