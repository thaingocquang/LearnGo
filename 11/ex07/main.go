package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Try again")
	} else {
		year, err := strconv.Atoi(os.Args[1])
		if err != nil || year < 0 {
			fmt.Println("Try again")
		} else {
			if year%400 == 0 {
				fmt.Println(year, "is a leap year.")
				return
			}
			if year%100 == 0 {
				fmt.Println(year, "is not a leap year.")
				return
			}
			if year%4 == 0 {
				fmt.Println(year, "is a leap year.")
				return
			}
			fmt.Println(year, "is not a leap year.")
			return
		}
	}
}
