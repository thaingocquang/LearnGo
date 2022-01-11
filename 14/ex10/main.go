package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{"Harry Potter", "Sherlock Holmes", "Gone with the Wind", "Nobody's Boy"}
	if len(os.Args) != 2 {
		fmt.Println("Pls input a book name")
		return
	}
	bookName := strings.ToLower(os.Args[1])

	isExist := false

	for _, book := range books {
		if strings.Contains(strings.ToLower(book), bookName) {
			fmt.Println(book)
			isExist = true
		}
	}
	if !isExist {
		fmt.Println("No book matched")
	}

}
