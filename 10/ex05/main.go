package main

import "fmt"

func main() {
	const (
		width  int32 = 25
		height int32 = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}
