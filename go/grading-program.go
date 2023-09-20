package main

import (
	"fmt"
)

func main() {
	student_scores := map[string]int{
		"Harry":    81,
		"Ron":      78,
		"Hermione": 99,
		"Draco":    74,
		"Neville":  62,
	}

	student_grades := map[string]string{}

	for key, value := range student_scores {
		var grade string
		if value > 90 {
			grade = "Outstanding"
		} else if value > 80 {
			grade = "Exceeds Expectations"
		} else if value > 70 {
			grade = "Acceptable"
		} else {
			grade = "Fail"
		}
		student_grades[key] = grade
	}

	fmt.Println(student_grades)
}
