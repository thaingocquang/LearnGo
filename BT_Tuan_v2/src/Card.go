package main

type Card struct {
	rank int
	suit int // co, ro, chuon, bich
}

// so sanh suit 2 card
func (c Card) IsHigher(c2 Card) bool {
	if c.suit < c2.suit {
		return true
	} else if c.suit > c2.suit {
		return false
	} else if c.suit == c2.suit {
		// truong hop trung con A
		realRank1 := c.rank
		realRank2 := c2.rank
		if realRank1 == 1 {
			realRank1 += 99
		} else if realRank2 == 1 {
			realRank2 += 99
		}
		if realRank1 > realRank2 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
