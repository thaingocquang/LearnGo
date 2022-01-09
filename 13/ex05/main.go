package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input 2 numbers")
	} else {
		num1, err1 := strconv.Atoi(os.Args[1])
		num2, err2 := strconv.Atoi(os.Args[2])
		if err1 != nil || err2 != nil || num1 > num2 {
			fmt.Println("Wrong number!")
		} else {
			sum := 0
			i := num1
			for {
				if i > num2 {
					break
				}
				if i%2 != 0 {
					i++
					continue
				}
				sum += i
				fmt.Print(i)
				if i < num2 {
					fmt.Print(" + ")
				}
				i++
			}

			fmt.Printf(" = %d\n", sum)
		}
	}
}
