package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) != 1 {
		fmt.Println("Give me a letter")
	} else {
		l := os.Args[1]
		if l == "a" || l == "i" || l == "u" || l == "e" || l == "o" {
			fmt.Println(`"` + l + `" is a vowel.`)
		} else if l == "y" || l == "w" {
			fmt.Println(`"` + l + `" is sometimes a vowel, sometimes not.`)
		} else {
			fmt.Println(`"` + l + `" is a consonant.`)
		}
	}

}
