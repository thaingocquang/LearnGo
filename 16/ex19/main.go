package main

import "fmt"

func main() {
	// var games []int
	// games := []string{}
	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Println("game's len and cap:", len(games), cap(games))
	games := []string{"pacman", "mario", "tetris", "doom"}
	for i := 0; i < len(games); i++ {
		slice := games[:i]
		fmt.Println("len", len(slice), "cap", cap(slice))
	}
	zero := games[:0]
	fmt.Println("zero slice len", len(zero), " zero slice cap", cap(zero))
	zero = append(zero, "new element")
	fmt.Println("zero slice len", len(zero), " zero slice cap", cap(zero))
	fmt.Println()
	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Println("zero slice len", len(zero), " zero slice cap", cap(zero))
	}
	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}
	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)
	games = games[:cap(games)]
	fmt.Println(games)
}
