package main

import (
	"math/rand"
	"time"
)

// ...
const totalSuit = 4
const cardSuit = 13

// DeckCard...
type DeckCard struct {
	deckCard []Card
}

// initial deck card
func (dc *DeckCard) init() {
	dc.deckCard = make([]Card, 0)
	for i := 1; i <= totalSuit; i++ {
		dc.deckCard = append(dc.deckCard, Card{1, 14, i})
		for j := 2; j <= cardSuit; j++ {
			dc.deckCard = append(dc.deckCard, Card{j, j, i})
		}
	}
}

// xao bai
func (dc *DeckCard) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(dc.deckCard), func(i, j int) { dc.deckCard[i], dc.deckCard[j] = dc.deckCard[j], dc.deckCard[i] })
}
