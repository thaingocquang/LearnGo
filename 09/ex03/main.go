package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	int8Val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	int16Val, _ := strconv.ParseInt(os.Args[2], 10, 16)
	int32Val, _ := strconv.ParseInt(os.Args[3], 10, 32)
	int64Val, _ := strconv.ParseInt(os.Args[4], 10, 64)
	bitVal, _ := strconv.ParseInt(os.Args[5], 2, 64)
	fmt.Println("int8 value is:", int8Val)
	fmt.Println("int8 value is:", int16Val)
	fmt.Println("int8 value is:", int32Val)
	fmt.Println("int8 value is:", int64Val)
	fmt.Println(os.Args[5], "is", bitVal)
}
