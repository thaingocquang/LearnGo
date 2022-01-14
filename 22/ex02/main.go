package main

import "fmt"

func main() {
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}
	products := map[int]bool{
		617841573: true,
		879401371: false,
		576872813: true,
	}
	multiPhones := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}
	basket := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
		102: {},
	}

	fmt.Println("dulin's phone number: ", phones["dulin"])
	if products[879401371] {
		fmt.Println("Product ID 879401371 available")
	} else {
		fmt.Println("Product ID 879401371 not available")
	}
	fmt.Println("Greco's second phone number is: ", multiPhones["greco"][1])
	fmt.Println("576872813 the customer 101 is going to buy", basket[101][576872813])
}
