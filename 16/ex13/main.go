package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	splitArr := strings.Fields(data)
	var nums []int
	for _, v := range splitArr {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	fmt.Println("nums", nums)
	fmt.Println("even", nums[:3])
	fmt.Println("odds", nums[3:])
	fmt.Println("middle", nums[2:4])
	fmt.Println("first 2", nums[:2])
	fmt.Println("last 2", nums[4:])
	fmt.Println("evens last 1", nums[2:3])
	fmt.Println("odds last 2", nums[4:])
}
