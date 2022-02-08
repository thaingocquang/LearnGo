package card

// Card...
type Card struct {
	Rank     int
	RealRank int
	Suit     int // co, ro, chuon, bich
}

// IsHigher so sanh 2 card
func (c Card) IsHigher(c2 Card) bool {
	if c.RealRank == c2.RealRank {
		return c.Suit < c2.Suit
	}
	return c.RealRank > c2.RealRank
}
