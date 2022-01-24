package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	var nilPointer *computer

	if nilPointer == nil {
		fmt.Println("nilPointer is nil")
	}

	applePointer := &computer{brand: "apple"}
	newApplePointer := applePointer

	if applePointer == newApplePointer {
		fmt.Println("2 pointer equals")
	}

	sonyPointer := &computer{brand: "sony"}

	if applePointer != sonyPointer {
		fmt.Println("2 pointer not equals")
	}

	// appleVal := *applePointer
	fmt.Printf("apple pointer variables address: %p, %p", &applePointer, applePointer)

	fmt.Println("=========")
	change(applePointer, "thai ngoc quang")
	fmt.Println(*applePointer)
}

func change(c *computer, brand string) {
	c.brand = brand
}
