package main

import "fmt"

func main() {
	var intVar int
	var intVar8 int8
	var intVar16 int16
	var intVar32 int32
	var intVar64 int64
	var floatVar32 float32
	var floatVar64 float64
	var complexVar64 complex64
	var complexVar128 complex128
	var boolVar bool
	var stringVar string
	var runeVar rune
	var byteVar byte

	fmt.Println(intVar, intVar8, intVar16, intVar32, intVar64,
		floatVar32, floatVar64,
		complexVar64, complexVar128,
		boolVar, stringVar, runeVar, byteVar)
}
