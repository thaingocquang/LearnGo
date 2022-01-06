package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t, _ := time.ParseDuration("1h30m")
	commandLineVar, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t = t * time.Duration(commandLineVar)
	fmt.Println(t)
}
