package main

import (
	"BT_Tuan_Package/src/card"
	"BT_Tuan_Package/src/player"
	"fmt"
)

func main() {
	var (
		totalPlayers int
		deckCard     card.DeckCard
		listPlayers  player.ListPlayers
	)
	// init bo bai
	deckCard.Init()
	for {
		fmt.Print("Nhap so nguoi choi:(2<=n<=17) ")
		fmt.Scanln(&totalPlayers)
		// xao bai
		deckCard.Shuffle()
		// init list players, chia bai, tinh kq
		listPlayers.Init(totalPlayers, deckCard)
		// hien thi bai nguoi choi va ket qua
		listPlayers.DisplayCards()
	}

}
