// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

func main() {
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)

	lyric = append([]string{"yesterday"}, lyric...)

	const N, M = 8, 13

	// yesterday all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay oh I believe in yesterday
	// yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
	fmt.Println(lyric[N:M])
	lyric = append(lyric, lyric[N:M]...)

	lyric = append(lyric[:N], lyric[M:]...)

	fmt.Printf("%s\n", lyric)
}
