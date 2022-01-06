package main

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Feet   float64
	Meters float64
)

func main() {
	var (
		feet   Feet
		meters Meters
	)
	arg := os.Args[1]
	feetVal, _ := strconv.ParseFloat(arg, 64)
	feet = Feet(feetVal)

	meters = Meters(feet * 0.3048)
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
