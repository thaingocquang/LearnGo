package main

import "fmt"

func main() {

	_, b := multi()

	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
