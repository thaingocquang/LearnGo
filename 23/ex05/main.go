package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wMap := make(map[string]bool)
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		if wMap[w] {
			fmt.Println("TWICE!!")
			return
		}
		wMap[w] = true
	}
}
