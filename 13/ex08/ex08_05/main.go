package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	pcRandTimes = 5
	usage       = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) <= 1 {
		fmt.Printf(usage, pcRandTimes)
		return
	}
	var isShownum = false
	if os.Args[1] == "-v" {
		isShownum = true
	}
	userGuessNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if userGuessNum <= 0 {
		fmt.Println("Pls guess a positive number")
		return
	}
	for i := 1; i <= pcRandTimes; i++ {
		var n int
		if userGuessNum <= 10 {
			n = rand.Intn(10) + 1
		} else {
			n = rand.Intn(userGuessNum) + 1
		}
		if isShownum {
			fmt.Printf("%d ", n)
		}
		if n == userGuessNum {
			if i == 1 {
				fmt.Println("FIRST TIME WINNER!!!")
			} else {
				fmt.Println("YOU WON!!")
			}
			return
		}
	}
	fmt.Println("YOU LOST.")
}
