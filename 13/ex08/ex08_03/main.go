package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns int = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 3 {
		fmt.Println("Pls guess 2 numbers")
		return
	}
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Not a number")
		return
	}
	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Pls guess positive number")
		return
	}
	greaterNum := num1
	if num2 > num1 {
		greaterNum = num2
	}
	for turn := 0; turn < maxTurns; turn++ {
		pcRandNum := rand.Intn(greaterNum) + 1
		if pcRandNum == num1 || pcRandNum == num2 {
			fmt.Println("YOU WON!!!")
			return
		}
	}
	fmt.Println("YOU LOST.")
}
