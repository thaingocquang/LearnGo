package main

import (
	"fmt"
)

func main() {

	var (
		nPlayers    int
		deckCard    DeckCard
		listPlayers ListPlayers
	)
	// init bo bai
	deckCard.init()
	for {
		fmt.Print("Nhap so nguoi choi:")
		fmt.Scanln(&nPlayers)
		// xao bai
		deckCard.Shuffle()
		// init list players va chia bai
		listPlayers.init(nPlayers, deckCard)
		// hien thi bai nguoi choi va ket qua
		listPlayers.DisplayCards()
	}

}
