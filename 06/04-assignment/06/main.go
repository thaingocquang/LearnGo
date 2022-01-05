package main

import "fmt"

func main() {

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet = "Air is good on Mars"
	isTrue = true
	temp = 19.5

	fmt.Println(planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")
}
