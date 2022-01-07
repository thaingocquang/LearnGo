package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
	} else {
		richter, err := strconv.ParseFloat(os.Args[1], 32)
		if err != nil {
			fmt.Println("I couldn't get that, sorry.")
		} else {
			var description string
			switch {
			case richter < 2:
				description = "micro"
			case richter >= 2 && richter < 3:
				description = "very minor"
			case richter >= 3 && richter < 4:
				description = "minor"
			case richter >= 4 && richter < 5:
				description = "light"
			case richter >= 5 && richter < 6:
				description = "moderate"
			case richter >= 6 && richter < 7:
				description = "strong"
			case richter >= 7 && richter < 8:
				description = "major"
			case richter >= 8 && richter < 10:
				description = "great"
			default:
				description = "massive"
			}
			fmt.Printf("%.2f is %s\n", richter, description)
		}
	}
}
