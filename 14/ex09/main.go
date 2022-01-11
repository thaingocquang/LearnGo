package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		EUR = iota
		GBP
		JPY
	)
	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}
	if len(os.Args) != 2 {
		fmt.Println("Please input amount to be converted")
		return
	}
	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Wrong amount. Not a number")
		return
	}
	fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
	fmt.Printf("%.2f GBP is %.2f GBP\n", amount, rates[GBP]*amount)
	fmt.Printf("%.2f JPY is %.2f JPY\n", amount, rates[JPY]*amount)
}
