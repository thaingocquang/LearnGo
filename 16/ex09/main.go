package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onion", "extra cheese")

	fmt.Printf("toppings: %s\n", pizza)
}
