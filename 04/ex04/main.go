package main

import "fmt"
import "runtime"

func main() {
	fmt.Print(runtime.Version())
	// output: go1.17.5
}
