package main

import "fmt"

func main() {
	fmt.Println("test")
	fmt.Println("test2")
	fmt.Println("test3")
	// goland auto format code
	// - before:
	// fmt.Println("test"); fmt.Println("test2"); fmt.Println("test3")
	// - after:
	// fmt.Println("test")
	// fmt.Println("test2")
	// fmt.Println("test3")
}
