package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [...]string{"Harry Poster", "sherlock holmes", "Gone with the Wind"}
	lower, upper := books, books
	for i := range books {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(upper[i])
	}
	fmt.Println(books)
	fmt.Println(upper)
	fmt.Println(lower)
}
