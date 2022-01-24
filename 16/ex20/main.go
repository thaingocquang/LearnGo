package main

import "fmt"

func main() {
	var (
		nums   []int
		oldCap float64
	)

	for i := 0; i < 1e7; i++ {
		nums = append(nums, 1)
		cap := float64(cap(nums))
		if oldCap == 0 || cap != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), cap, cap/oldCap)
		}
		oldCap = cap
		nums = append(nums, 1)
	}
}
