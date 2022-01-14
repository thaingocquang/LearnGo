package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%q\n", ships)
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	}
	if len(args) > 2 {
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	}
	var index1, index2 int
	if len(args) == 1 {
		n, _ := strconv.Atoi()
	}
	if args[0] < 0 || args[1] < 0 || args[1] > 1 || args[0] > 1 {

	}
}
