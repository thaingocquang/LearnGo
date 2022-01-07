package main

import "fmt"

func main() {
	// NOTE : You should remove all the initializers below
	//        first. Then use iota to fix it.
	const (
		Spring = 3 + iota*3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
