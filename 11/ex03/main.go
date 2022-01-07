package main

import (
	"fmt"
	"os"
)

func main() {
	argLen := len(os.Args) - 1

	if argLen == 0 {
		fmt.Println("Give me args")
	}
	if argLen == 1 {
		fmt.Println(`There is one: "`, os.Args[1], `"`)
	}
	if argLen == 2 {
		fmt.Println(`There is two: "`, os.Args[1], os.Args[2], `"`)
	}
	if argLen == 5 {
		fmt.Println("There are 5 arguments")
	}
}
