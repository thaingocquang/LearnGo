package main

import "fmt"

func main() {
	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	// ----------------------------------------
	// FIX THE CODE BELOW:
	// You should solve it only by using conversions.
	// Do not change the code in any other way.

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	// ----------------------------------------
	// DONT TOUCH THIS
	fmt.Println(celsius, fahr)
}
