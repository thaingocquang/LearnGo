package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	if len(os.Args) <= 1 {
		fmt.Println("Pls input a house name")
		return
	}

	house, students := os.Args[1], houses[os.Args[1]]
	if students == nil {
		fmt.Printf("Sorry, i dont know anything about %q", house)
		return
	}

	copyStudents := append([]string(nil), students...)
	sort.Strings(copyStudents)
	for _, v := range copyStudents {
		fmt.Println(v)
	}
}
