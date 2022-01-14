package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a few number")
		return
	}
	var nums []int
	for _, v := range args {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
