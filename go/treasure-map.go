package main

import (
	"fmt"
	"strconv"
)

func main() {
	row1 := []string{"⬜️", "️⬜️", "️⬜️"}
	row2 := []string{"⬜️", "⬜️", "️⬜️"}
	row3 := []string{"⬜️️", "⬜️️", "⬜️️"}

	matrix := [][]string{
		row1,
		row2,
		row3,
	}

	fmt.Print("Where do you want to put the treasure? ")
	var position string
	fmt.Scanln(&position)

  column_idx, err := strconv.Atoi(string(position[0]))
  if err != nil {
    panic(err)
  }

  row_idx, err := strconv.Atoi(string(position[1]))
  if err != nil {
    panic(err)
  }

  matrix[row_idx-1][column_idx-1] = "X"

  fmt.Println(matrix)
}
