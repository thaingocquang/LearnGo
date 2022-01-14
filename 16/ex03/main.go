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

	names = []string{"Cao", "Danh", "Thai"}
	distances = []int{15, 23, 44}
	data = []byte{'t', 'n', 'q'}
	ratios = []float64{3.14, 6.28}
	alives = []bool{true, false}

	fmt.Printf("%T - %d - %t\n", names, len(names), names == nil)
	fmt.Printf("%T - %d - %t\n", distances, len(distances), distances == nil)
	fmt.Printf("%T - %d - %t\n", data, len(data), data == nil)
	fmt.Printf("%T - %d - %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("%T - %d - %t\n", alives, len(alives), alives == nil)
}
