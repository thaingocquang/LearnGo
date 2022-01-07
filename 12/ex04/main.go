package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = `[command] [string]
Available commands: lower, upper and title`

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	command, str := os.Args[1], os.Args[2]
	switch command {
	case "lower":
		fmt.Println(strings.ToLower(str))
	case "upper":
		fmt.Println(strings.ToUpper(str))
	case "title":
		fmt.Println(strings.Title(str))
	default:
		fmt.Printf("Unknown command: %q\n", command)
	}
}
