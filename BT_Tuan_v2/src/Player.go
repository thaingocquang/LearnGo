package main

// Player...
type Player struct {
	id       int
	cards    []Card
	sumCards int
	highCard Card
}

// CalculateSumCards...
func (p *Player) CalculateSumCards() {
	for i := 0; i < threeCard; i++ {
		if p.cards[i].rank >= 10 {
			p.sumCards += 0
		} else {
			p.sumCards += p.cards[i].rank
		}
	}
	p.sumCards %= 10
}

// CalculateHighCard find the card which has highest rank and suit
func (p *Player) CalculateHighCard() {
	p.highCard = p.cards[0]
	for i := 1; i < threeCard; i++ {
		if p.cards[i].IsHigher(p.highCard) {
			p.highCard = p.cards[i]
		}
	}
}
