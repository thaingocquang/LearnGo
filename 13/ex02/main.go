package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Print(i)
		if i < 10 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d", sum)
}
