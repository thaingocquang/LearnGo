package card

// Card...
type Card struct {
	Rank     int
	RealRank int
	Suit     int // co, ro, chuon, bich
}

// so sanh 2 card
func (c Card) IsHigher(c2 Card) bool {
	if c.RealRank > c2.RealRank {
		return true
	} else if c.RealRank < c2.RealRank {
		return false
	} else {
		if c.Suit < c2.Suit {
			return true
		} else {
			return false
		}
	}
}
