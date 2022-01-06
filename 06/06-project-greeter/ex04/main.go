package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		_len  = len(os.Args) - 1
		name1 = os.Args[1]
		name2 = os.Args[2]
		name3 = os.Args[3]
	)

	fmt.Println("There are ", _len, " people!")
	fmt.Println("Hello great " + name1 + " !")
	fmt.Println("Hello great " + name2 + " !")
	fmt.Println("Hello great " + name3 + " !")
	fmt.Println("Nice to meet you all.")
}
