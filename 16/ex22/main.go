package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	var names = make([]string, 5)
	s.Show("1st step", names)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("2st step", names)
	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("3rd step", names)
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)] 
	fmt.Println(names)
}
