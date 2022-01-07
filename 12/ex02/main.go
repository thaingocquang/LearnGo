package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
	} else {
		description := os.Args[1]
		switch description {
		case "micro":
			fmt.Println(description + "'s richter scale is less than 2.0")
		case "very minor":
			fmt.Println(description + "'s richter scale is 2.0 to less than 3.0 ")
		case "minor":
			fmt.Println(description + "'s richter scale is 3.0 to less than 4.0 ")
		case "light":
			fmt.Println(description + "'s richter scale is 4.0 to less than 5.0  ")
		case "moderate":
			fmt.Println(description + "'s richter scale is 5.0 to less than 6.0")
		case "strong":
			fmt.Println(description + "'s richter scale is 6.0 to less than 7.0")
		case "major":
			fmt.Println(description + "'s richter scale is 7.0 to less than 8.0")
		case "great":
			fmt.Println(description + "'s richter scale is 8.0 to less than 10.0")
		case "massive":
			fmt.Println(description + "'s richter scale is 10.0 or more  ")
		default:
			fmt.Println(description + "'s richter scale is unknown")
		}
	}
}
