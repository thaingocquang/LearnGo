package main

import (
	"fmt"
	"os"
)

func main() {
	// uncomment the code below
	name := os.Args[1]

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
}
