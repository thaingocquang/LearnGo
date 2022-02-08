package main

// Card...
type Card struct {
	rank     int
	realRank int
	suit     int // co, ro, chuon, bich
}

// so sanh 2 card
func (c Card) IsHigher(c2 Card) bool {
	if c.realRank > c2.realRank {
		return true
	} else if c.realRank < c2.realRank {
		return false
	} else {
		if c.suit < c2.suit {
			return true
		} else {
			return false
		}
	}
}
