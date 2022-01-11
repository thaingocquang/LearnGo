package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Pls input any query word")
		return
	}
	arrCorpus := strings.Fields(corpus)
	for _, v := range args {
		v = strings.ToLower(v)
		if v == "and" || v == "or" || v == "was" || v == "the" || v == "since" || v == "very" {
			continue
		}
		for _i, _v := range arrCorpus {
			if v == _v {
				fmt.Printf("%d - %q\n", _i+1, v)
			}
		}
	}

}
