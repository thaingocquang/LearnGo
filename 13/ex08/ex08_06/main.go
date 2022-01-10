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
	arg := os.Args[1:]
	if len(arg) < 1 {
		fmt.Printf(usage, pcRandTimes)
		return
	}
	var isShowNum bool = false
	if arg[0] != "-v" {
		isShowNum = true
	}
	userGuessNum, err := strconv.Atoi(arg[len(arg)-1])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if userGuessNum <= 0 {
		fmt.Println("Pls guess a positive number")
		return
	}
	var turn int
	if userGuessNum <= 10 {
		turn = 5
	} else if userGuessNum <= 100 {
		turn = 11
	} else {
		turn = 30
	}
	for i := 1; i <= turn; i++ {
		var n int
		if turn == 5 {
			n = rand.Intn(5) + 1
		} else if turn == 11 {
			n = rand.Intn(100) + 1
		} else {
			n = rand.Intn(1000) + 1
		}
		if isShowNum {
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
