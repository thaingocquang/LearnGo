package main

import (
	"fmt"
	"os"
)

func main() {
	message := "My name is %s and my last name is %s"
	fmt.Printf(message, os.Args[1], os.Args[2])
}
