package main

import "fmt"

func main() {
	var message string

	fmt.Printf("My first name is %s and my last name is %s\n", "Quang", "Thai Ngoc")

	message = "My first name is %s and my last name is %s\n"
	fmt.Printf(message, "Quang", "Thai Ngoc")
}
