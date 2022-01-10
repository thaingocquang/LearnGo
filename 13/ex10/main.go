package main

import (
	"fmt"
	"os"
	"strconv"
)

func isprime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}
	var i int = 5
	var w int = 2

	for (i * i) <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Pls input at least 1 number")
	}
	listNum := os.Args[1:]
	for _, num := range listNum {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		if isprime(intNum) {
			fmt.Printf("%d ", intNum)
		}
	}
}
