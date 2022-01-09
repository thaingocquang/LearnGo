package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	win := [3]string{"🎉  YOU WIN!", "🎉  YOU'RE AWESOME!", "🎉  PERFECT!"}
	lose := [2]string{"☠️  YOU LOST...", "☠️  JUST A BAD LUCK..."}
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}
	if guess <= 0 {
		fmt.Println("Pls guess a positive number")
		return
	}
	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(guess) + 1
		if n == guess {
			index := rand.Intn(3)
			fmt.Println(win[index])
			return
		}
	}
	index := rand.Intn(2)
	fmt.Println(lose[index])

}
