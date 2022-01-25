package player

import "BT_Tuan_Package/src/card"

// Player...
type Player struct {
	Id       int
	Cards    []card.Card
	SumCards int
	HighCard card.Card
}

// CalculateSumCards...
func (p *Player) CalculateSumCards() {
	for i := 0; i < threeCard; i++ {
		if p.Cards[i].Rank >= 10 {
			p.SumCards += 0
		} else {
			p.SumCards += p.Cards[i].Rank
		}
	}
	p.SumCards %= 10
}

// CalculateHighCard find the card which has highest rank and suit
func (p *Player) CalculateHighCard() {
	p.HighCard = p.Cards[0]
	for i := 1; i < threeCard; i++ {
		if p.Cards[i].IsHigher(p.HighCard) {
			p.HighCard = p.Cards[i]
		}
	}
}
