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
	guess, _ := strconv.Atoi(os.Args[1])
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("The first time winner!!")
		} else {
			fmt.Println("You won")
		}
		return
	}
	fmt.Println("You lost... Try again???")
}
