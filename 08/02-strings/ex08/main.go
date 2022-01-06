package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	name := "inanç           "
	name = strings.TrimRight(name, " ")
	_len := utf8.RuneCountInString(name)
	fmt.Println(_len)
}
