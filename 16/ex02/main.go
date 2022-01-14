package main

import "fmt"

func main() {
	var (
		names     []string
		distances []int
		data      []byte
		ratios    []float64
		alives    []bool
	)

	names = []string{}
	distances = []int{}
	data = []byte{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("%T - %d - %t\n", names, len(names), names == nil)
	fmt.Printf("%T - %d - %t\n", distances, len(distances), distances == nil)
	fmt.Printf("%T - %d - %t\n", data, len(data), data == nil)
	fmt.Printf("%T - %d - %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("%T - %d - %t\n", alives, len(alives), alives == nil)
}
