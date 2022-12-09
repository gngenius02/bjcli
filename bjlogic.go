package main

type BJ struct {
	playerCards []int
	lenPC       int
	handValue   int
	handSoft    bool
	dealerCard  int
	handCount   int
}

func NewBJ(ph []int, dc int, hc int) *BJ {
	b := BJ{playerCards: ph, dealerCard: dc, handCount: hc}
	b.calcHandTotal()
	b.lenPC = len(b.playerCards)
	return &b
}

func (b *BJ) calcHandTotal() {
	hasAces := false
	for _, v := range b.playerCards {
		b.handValue += v
		if v == 1 {
			hasAces = true
		}
	}
	if b.handValue <= 11 && hasAces {
		b.handValue += 10
		b.handSoft = true
	}
}

func includes(in []int, i int) bool {
	for _, v := range in {
		if v == i {
			return true
		}
	}
	return false
}

func (b *BJ) BlackjackLogic() string {
	if b.playerCards[0] == b.playerCards[1] && b.lenPC == 2 && b.handCount < 2 {
		switch {
		case b.playerCards[0] == 1 || b.playerCards[0] == 8:
			return "split"
		case b.playerCards[0] == 4 && includes([]int{5, 6}, b.dealerCard):
			return "split"
		case b.playerCards[0] == 6 && includes([]int{2, 3, 4, 5, 6}, b.dealerCard):
			return "split"
		case b.playerCards[0] == 9 && !includes([]int{1, 7, 10}, b.dealerCard):
			return "split"
		case includes([]int{2, 3, 7}, b.playerCards[0]) && !includes([]int{1, 8, 9, 10}, b.dealerCard):
			return "split"
		}
	}
	if b.lenPC == 2 {
		switch {
		case b.handValue == 9 && includes([]int{3, 4, 5, 6}, b.dealerCard):
			return "double"
		case b.handValue == 10 && b.dealerCard < 10 && b.dealerCard != 1:
			return "double"
		case b.handValue == 11 && b.dealerCard > 1:
			return "double"
		case includes(b.playerCards, 1):
			switch {
			case (includes(b.playerCards, 6) || includes(b.playerCards, 7)) && includes([]int{3, 4, 5, 6}, b.dealerCard):
				return "double"
			case includes(b.playerCards, 2) && b.dealerCard == 6:
				return "double"
			case includes(b.playerCards, 5) && includes([]int{4, 5, 6}, b.dealerCard):
				return "double"
			case (includes(b.playerCards, 3) || includes(b.playerCards, 4)) && includes([]int{5, 6}, b.dealerCard):
				return "double"
			}
		}
	}
	if b.handSoft {
		switch {
		case b.handValue < 18:
			return "hit"
		case b.handValue == 18 && (b.dealerCard > 8 || b.dealerCard == 1):
			return "hit"
		default:
			return "stand"
		}
	}
	switch {
	case b.handValue <= 11:
		return "hit"
	case b.handValue == 12 && includes([]int{2, 3}, b.dealerCard):
		return "hit"
	case b.dealerCard < 7 && b.dealerCard != 1:
		return "stand"
	case b.handValue < 17:
		return "hit"
	default:
		return "stand"
	}
}
