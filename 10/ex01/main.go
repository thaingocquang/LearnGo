package main

import "fmt"

func main() {
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)
	fmt.Println("There are", 2*weekDays*minsPerDay, "minutes in 2 weeks.")
}
