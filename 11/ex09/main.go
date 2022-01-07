package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	} else {
		return false
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	_m := os.Args[1]
	m := strings.ToLower(_m)
	t := time.Now()

	if m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		fmt.Printf("%q has 30 days.\n", _m)
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		fmt.Printf("%q has 31 days.\n", _m)
	} else if m == "february" {
		if isLeapYear(t.Year()) {
			fmt.Printf("%q has 29 days.\n", _m)
		} else {
			fmt.Printf("%q has 28 days.\n", _m)
		}
	} else {
		fmt.Printf("%q is not a month.\n", _m)
	}

}
