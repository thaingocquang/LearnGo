package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}
	for i := range names {
		e := names[i]
		fmt.Printf("%10s %10s %10s\n", e[0], e[1], e[2])
		if i == 0 {
			fmt.Println(strings.Repeat("=", 35))
		}
	}
}
