package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 0 || len(os.Args) > 6 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}
	var (
		sum   float64
		nums  [5]float64
		total float64
	)
	args := os.Args[1:]
	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		total++
		nums[i] = n
		sum += n
	}
	fmt.Println("Your numbers: ", nums)
	fmt.Println("Average: ", sum/total)
}
