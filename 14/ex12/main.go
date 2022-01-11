package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Pls input up to 5 number")
		return
	}
	if len(args) > 5 {
		fmt.Println("Maximum 5 numbers allowed")
		return
	}
	var _args [5]float64
	for i, v := range args {
		num, _ := strconv.ParseFloat(v, 64)
		_args[i] = num
	}
	n := len(_args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if _args[j] > _args[j+1] {
				_args[j], _args[j+1] = _args[j+1], _args[j]
			}
		}
	}
	fmt.Println(_args)
}
