package main

import "fmt"

func main() {
	isSphere, radius := true, 200
	if isSphere == true && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}
