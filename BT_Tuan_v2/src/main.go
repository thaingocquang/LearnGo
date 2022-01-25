package main

import (
	"fmt"
)

func main() {
	var (
		totalPlayers int
		deckCard     DeckCard
		listPlayers  ListPlayers
	)
	// init bo bai
	deckCard.init()
	for {
		fmt.Print("Nhap so nguoi choi:(2<=n<=17) ")
		fmt.Scanln(&totalPlayers)
		// xao bai
		deckCard.Shuffle()
		// init list players, chia bai, tinh kq
		listPlayers.init(totalPlayers, deckCard)
		// hien thi bai nguoi choi va ket qua
		listPlayers.DisplayCards()
	}

}
