package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usageMess       = "Usage: [op=*/+-] [size]"
	sizeMissingMess = "Size is missing"
	invalidOpMess   = `Invalid operator.
Valid ops one of: */+-`
	wrongSizeMess = "Wrong size."
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usageMess)
	} else {
		op := os.Args[1]
		if strings.IndexAny(op, "%*/+-") == -1 {
			fmt.Println(invalidOpMess)
		} else {
			size, err := strconv.Atoi(os.Args[2])
			if err != nil || size < 0 {
				fmt.Println(wrongSizeMess)
			} else {
				fmt.Printf("%5s", op)
				for i := 0; i <= size; i++ {
					fmt.Printf("%5d", i)
				}
				fmt.Println()
				for i := 0; i <= size; i++ {
					fmt.Printf("%5d", i)
					for j := 0; j <= size; j++ {
						var res int
						if op == "+" {
							res = i + j
						}
						if op == "-" {
							res = i - j
						}
						if op == "*" {
							res = i * j
						}
						if op == "/" {
							if j != 0 {
								res = i / j
							}
						}
						if op == "%" {
							if j != 0 {
								res = i % j
							}
						}
						fmt.Printf("%5d", res)
					}
					fmt.Println()
				}
			}
		}
	}
}
