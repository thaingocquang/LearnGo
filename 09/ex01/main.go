package main

import "fmt"

func main() {
	// an english letter (search web for: ascii control code)
	var a byte = 'a'
	fmt.Println("an english letter:", a)

	// a non-english letter (search web for: unicode codepoint)
	var nel rune = 'Æ'
	fmt.Println(nel)

	// a year in 4 digits like 2040
	var year uint16 = 2022
	fmt.Println(year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 12
	fmt.Println(month)

	// the speed of the light
	var sol uint32 = 186000
	fmt.Println("speed of the light is", sol, "mile per second")

	// angle of a circle
	var angle float32 = 90.5
	fmt.Println(angle)

	// PI number: 3.141592653589793
	var pi float64 = 3.141592653589793
	fmt.Println(pi)
}
