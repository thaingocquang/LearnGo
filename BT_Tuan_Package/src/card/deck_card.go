package card

import (
	"math/rand"
	"time"
)

// ...
const totalSuit = 4
const cardSuit = 13

// DeckCard...
type DeckCard struct {
	DeckCard []Card
}

// Init khoi tao deck card
func (dc *DeckCard) Init() {
	dc.DeckCard = make([]Card, 0)
	for i := 1; i <= totalSuit; i++ {
		dc.DeckCard = append(dc.DeckCard, Card{1, 99, i}) // khoi tao con xi co rank 1 va real rank 99
		for j := 2; j <= cardSuit; j++ {
			dc.DeckCard = append(dc.DeckCard, Card{j, j, i})
		}
	}
}

// xao bai
func (dc *DeckCard) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dc.DeckCard), func(i, j int) { dc.DeckCard[i], dc.DeckCard[j] = dc.DeckCard[j], dc.DeckCard[i] })
}
