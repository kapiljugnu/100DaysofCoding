package main

import (
	"fmt"
)

func is_leap(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func days_in_month(year, month int) int {
	month_days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 && is_leap(year) {
		return 29
	}
	return month_days[month-1]
}

func main() {
	fmt.Print("Enter a year: ")
	var year int
	fmt.Scanln(&year)
	fmt.Print("Enter a month: ")
	var month int
	fmt.Scanln(&month)

	days := days_in_month(year, month)
	fmt.Println(days)
}
