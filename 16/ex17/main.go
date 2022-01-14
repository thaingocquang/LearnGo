package main

import "fmt"

func main() {
	mySt := st("quang")
	var intf I = &mySt
	intf.log("thaingocquang")
}

type I interface {
	log(string) bool
}

type st string

func (s *st) log(string) bool {
	fmt.Println(*s)
	return true
}
