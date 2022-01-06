package main

import "fmt"

func main() {
	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.
	var (
		width  uint16
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	fmt.Println("are they equal?", width == height)
}
