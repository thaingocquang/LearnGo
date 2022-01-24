package main

import (
	"math/rand"
	"time"
)

const nSuit = 4
const nCard1Suit = 13

type DeckCard struct {
	deckCard []Card
}

func (dc *DeckCard) init() {
	dc.deckCard = make([]Card, 0)
	for i := 1; i <= nSuit; i++ {
		for j := 1; j <= nCard1Suit; j++ {
			dc.deckCard = append(dc.deckCard, Card{j, i})
		}
	}
}
func (dc *DeckCard) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dc.deckCard), func(i, j int) { dc.deckCard[i], dc.deckCard[j] = dc.deckCard[j], dc.deckCard[i] })
}
