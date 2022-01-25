package player

import (
	"fmt"
	"sort"
	"time"

	"github.com/atomicgo/cursor"

	"BT_Tuan_Package/src/card"
)

// number of cards player have
const threeCard = 3

// rank, suit name array for display card
var (
	rankName = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suitName = []string{"cơ", "rô", "chuồn", "bích"}
)

// ListPlayers...
type ListPlayers struct {
	listPlayers []Player
}

// hien thi bai nguoi choi va ket qua
func (lp ListPlayers) DisplayCards() {
	fmt.Print("\033[H\033[2J")
	totalPlayers := len(lp.listPlayers)
	for i := 1; i <= totalPlayers; i++ {
		fmt.Println("Player", i, ":")
	}

	for k := 0; k < 3; k++ {
		cursor.Up(totalPlayers)
		if k == 0 {
			cursor.Right((k + 1) * 15)
		} else {
			cursor.Right(20)
		}
		for i := 0; i < totalPlayers; i++ {
			fmt.Printf("%2s%6s", rankName[lp.listPlayers[i].Cards[k].Rank-1], suitName[lp.listPlayers[i].Cards[k].Suit-1])
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
			if v.Id == winner.Id {
				fmt.Println("Player", v.Id, ": ", v.SumCards, "điểm", "=> Winner")
			} else {
				fmt.Println("Player", v.Id, ": ", v.SumCards, "điểm")
			}
		}
	} else {
		for _, v1 := range lp.listPlayers {
			if v1.Id == winner.Id {
				fmt.Println("Player", v1.Id, ": ", v1.SumCards, "điểm", "(", rankName[v1.HighCard.Rank-1], suitName[v1.HighCard.Suit-1], ")", "=> Winner")
			} else {
				isInEqualRank := false
				for _, v2 := range equalRank {
					if v1.Id == v2.Id {
						isInEqualRank = true
					}
				}
				if isInEqualRank {
					fmt.Println("Player", v1.Id, ": ", v1.SumCards, "điểm", "(", rankName[v1.HighCard.Rank-1], suitName[v1.HighCard.Suit-1], ")")
				} else {
					fmt.Println("Player", v1.Id, ": ", v1.SumCards, "điểm")
				}
			}
		}
	}
	fmt.Println()
}

// init player, chia bai, tinh kq
func (lp *ListPlayers) Init(totalPlayers int, shuffledDeckCard card.DeckCard) {
	lp.listPlayers = make([]Player, 0)
	var cardInx = 0
	for i := 1; i <= totalPlayers; i++ {
		player := Player{
			Id:       i,
			Cards:    make([]card.Card, threeCard),
			SumCards: 0,
			HighCard: card.Card{0, 0, 0},
		}
		for j := 0; j < threeCard; j++ {
			card := shuffledDeckCard.DeckCard[cardInx]
			player.Cards[j] = card
			cardInx++
		}
		player.CalculateSumCards()
		player.CalculateHighCard()
		lp.listPlayers = append(lp.listPlayers, player)
	}
}

// FindWinner find the player has highest score and suit
func (lpp ListPlayers) FindWinner() (winner Player, equalRank []Player) {
	// copy ra list player moi
	var lp ListPlayers
	lp.listPlayers = make([]Player, len(lpp.listPlayers))
	copy(lp.listPlayers, lpp.listPlayers)

	// sort list players theo diem
	sort.Slice(lp.listPlayers, func(i, j int) bool {
		return lp.listPlayers[i].SumCards > lp.listPlayers[j].SumCards
	})
	// tim slice player co diem cao nhat bang nhau
	k := 0
	for i := 1; i < len(lp.listPlayers); i++ {
		if lp.listPlayers[i].SumCards != lp.listPlayers[0].SumCards {
			break
		}
		k++
	}
	if k == 0 {
		return lp.listPlayers[0], nil
	} else {
		//sort list players
		sort.Slice(lp.listPlayers[0:k+1], func(i, j int) bool {
			return lp.listPlayers[i].HighCard.IsHigher(lp.listPlayers[j].HighCard)
		})
		return lp.listPlayers[0], lp.listPlayers[1 : k+1]
	}
}
