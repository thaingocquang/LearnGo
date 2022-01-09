package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the size of the table")
	} else {
		size, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Wrong size")
		} else {
			fmt.Printf("%5s", "X")
			for i := 0; i <= size; i++ {
				fmt.Printf("%5d", i)
			}
			fmt.Println()
			for i := 0; i <= size; i++ {
				fmt.Printf("%5d", i)
				for j := 0; j <= size; j++ {
					fmt.Printf("%5d", i*j)
				}
				fmt.Println()
			}
		}
	}
}
