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
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Try again")
		} else {
			if num%2 == 0 {
				if num%8 == 0 {
					fmt.Println(num, "is an even number and it's divisible by 8")
				} else {
					fmt.Println(num, "is an even number")
				}
			} else {
				fmt.Println(num, "is an odd number")
			}
		}
	}
}
