package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sliceNameA := strings.Split(namesA, ", ")
	sort.Strings(namesB)
	sort.Strings(sliceNameA)
	if len(namesB) == len(sliceNameA) {
		for i := range namesB {
			if namesB[i] != sliceNameA[i] {
				return
			}
		}
		fmt.Println("They are equal")
	}
}
