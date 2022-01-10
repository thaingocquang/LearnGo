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
	if len(os.Args) < 1 {
		fmt.Printf(usage, pcRandTimes)
	}
	isShowNum := false
	if os.Args[1] == "-v" {
		isShowNum = true
	}
	guessNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if guessNum <= 0 {
		fmt.Println("Pls guess a positive number.")
		return
	}
	for turn := 1; turn <= pcRandTimes; turn++ {
		n := rand.Intn(guessNum) + 1
		if isShowNum {
			fmt.Printf("%d ", n)
		}
		if n == guessNum {
			if turn == 1 {
				fmt.Println("FIRST TIME WINNER!!!")
			} else {
				fmt.Println("YOU WON!!")
			}
			return
		}
	}
	fmt.Println("YOU LOST.")
}
