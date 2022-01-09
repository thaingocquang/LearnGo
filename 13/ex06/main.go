package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// fmt.Print(i)

	arr := [4]string{"/", "\\", "|", "-"}

	for {
		i := rand.Intn(4)
		fmt.Printf("\r%s Please Wait. Processing...", arr[i])
		time.Sleep(time.Millisecond * 250)
	}

}
