package main

import "fmt"

func swap(a *float64, b *float64) {
	*a, *b = *b, *a
}

func main() {
	var a, b float64
	a = 3.1
	b = 6.1
	swap(&a, &b)
	fmt.Println(a, b)

}
