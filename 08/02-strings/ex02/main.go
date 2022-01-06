package main

import "fmt"

func main() {
	json := `
{
	"Item": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`
	fmt.Println(json)
}
