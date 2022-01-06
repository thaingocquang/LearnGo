package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]

	times := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", times)

	fmt.Println(s)
}
