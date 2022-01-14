package main

import (
	"fmt"
	"strconv"
	"strings"
)

func avg(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	rows := strings.Split(data, "\n")
	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	for _, row := range rows {
		r := strings.Split(row, separator)
		locs = append(locs, r[0])

		s, _ := strconv.Atoi(r[1])
		sizes = append(sizes, s)

		b, _ := strconv.Atoi(r[2])
		beds = append(beds, b)

		ba, _ := strconv.Atoi(r[3])
		baths = append(baths, ba)

		p, _ := strconv.Atoi(r[4])
		prices = append(prices, p)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-10s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 100))

	for i := range rows {
		fmt.Printf("%-10s", locs[i])
		fmt.Printf("%-10d", sizes[i])
		fmt.Printf("%-10d", beds[i])
		fmt.Printf("%-10d", baths[i])
		fmt.Printf("%-10d", prices[i])
		fmt.Println()
	}
	fmt.Printf("%s\n", strings.Repeat("=", 100))
	fmt.Printf("%-10.2s", "")
	fmt.Printf("%-10.2f", avg(sizes))
	fmt.Printf("%-10.2f", avg(beds))
	fmt.Printf("%-10.2f", avg(baths))
	fmt.Printf("%-10.2f", avg(prices))
}
