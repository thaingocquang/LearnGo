package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/atomicgo/cursor"
)

type ListPlayers struct {
	listPlayers []Player
}

func (lp ListPlayers) DisplayCards() {
	rankName := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suitName := []string{"cơ", "rô", "chuồn", "bích"}
	fmt.Print("\033[H\033[2J")
	nPlayers := len(lp.listPlayers)
	for i := 1; i <= nPlayers; i++ {
		fmt.Println("Player", i, ":")
	}

	for k := 0; k < 3; k++ {
		cursor.Up(nPlayers)
		if k == 0 {
			cursor.Right((k + 1) * 15)
		} else {
			cursor.Right(20)
		}
		for i := 0; i < nPlayers; i++ {

			fmt.Printf("%2s%6s", rankName[lp.listPlayers[i].cards[k].rank-1], suitName[lp.listPlayers[i].cards[k].suit-1])

			time.Sleep(time.Millisecond * 10)
			cursor.Down(1)
			cursor.Left(8)
		}
	}
	cursor.Bottom()
	cursor.Left(30)
	fmt.Println("KẾT QUẢ:")
	winner, equalRank := lp.FindWinner()
	if equalRank == nil {
		for _, v := range lp.listPlayers {
			if v.id == winner.id {
				fmt.Println("Player", v.id, ": ", v.sumCards, "điểm", "=> Winner")
			} else {
				fmt.Println("Player", v.id, ": ", v.sumCards, "điểm")
			}
		}
	} else {
		for _, v1 := range lp.listPlayers {
			if v1.id == winner.id {
				fmt.Println("Player", v1.id, ": ", v1.sumCards, "điểm", "(", rankName[v1.highCard.rank-1], suitName[v1.highCard.suit-1], ")", "=> Winner")
			} else {
				isInEqualRank := false
				for _, v2 := range equalRank {
					if v1.id == v2.id {
						isInEqualRank = true
					}
				}
				if isInEqualRank {
					fmt.Println("Player", v1.id, ": ", v1.sumCards, "điểm", "(", rankName[v1.highCard.rank-1], suitName[v1.highCard.suit-1], ")")
				} else {
					fmt.Println("Player", v1.id, ": ", v1.sumCards, "điểm")
				}
			}
		}
	}
	fmt.Println()
}

func (lp *ListPlayers) init(nPlayers int, shuffledDeckCard DeckCard) {
	lp.listPlayers = make([]Player, 0)
	var cardInx = 0
	for i := 1; i <= nPlayers; i++ {
		var player Player
		player.id = i
		player.cards = make([]Card, 3)
		player.highCard = Card{5, 5}
		for j := 0; j < 3; j++ {
			player.cards[j] = shuffledDeckCard.deckCard[cardInx]
			// Tinh diem sumCard cho player
			if player.cards[j].rank >= 10 {
				player.sumCards += 0
			} else {
				player.sumCards += player.cards[j].rank
			}
			// Tim highCard cho player
			if player.cards[j].IsHigher(player.highCard) {
				player.highCard = player.cards[j]
			}
			cardInx++
		}
		player.sumCards %= 10
		lp.listPlayers = append(lp.listPlayers, player)
	}
}

func (lpp ListPlayers) FindWinner() (winner Player, equalRank []Player) {
	// copy ra list player moi
	var lp ListPlayers
	lp.listPlayers = make([]Player, len(lpp.listPlayers))
	copy(lp.listPlayers, lpp.listPlayers)

	// sort list players theo diem
	sort.Slice(lp.listPlayers, func(i, j int) bool {
		return lp.listPlayers[i].sumCards > lp.listPlayers[j].sumCards
	})
	// tim slice player co diem cao nhat bang nhau
	k := 0
	for i := 1; i < len(lp.listPlayers); i++ {
		if lp.listPlayers[i].sumCards != lp.listPlayers[0].sumCards {
			break
		}
		k++
	}
	if k == 0 {
		return lp.listPlayers[0], nil
	} else {
		//sort list players theo suit
		sort.Slice(lp.listPlayers[0:k+1], func(i, j int) bool {
			return lp.listPlayers[i].highCard.IsHigher(lp.listPlayers[j].highCard)
		})
		return lp.listPlayers[0], lp.listPlayers[1 : k+1]
	}
}
