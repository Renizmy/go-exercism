package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1 string, card2 string, dealerCard string) string {

	scoreCard1 := ParseCard(card1)
	scoreCard2 := ParseCard(card2)
	scoreDealerCard := ParseCard(dealerCard)
	scoreHand := scoreCard1 + scoreCard2

	switch {
	case scoreCard1 == 11 && scoreCard2 == 11:
		return "P"

	case scoreHand == 21 && scoreDealerCard < 10:
		return "W"
	case scoreHand == 21 && scoreDealerCard > 9:
		return "S"

	case scoreHand > 16 && scoreHand < 21:
		return "S"

	case scoreHand > 11 && scoreHand <= 16 && scoreDealerCard > 6:
		return "H"
	case scoreHand > 11 && scoreHand <= 16 && scoreDealerCard < 7:
		return "S"
	case scoreHand < 12:
		return "H"

	default:
		return "H"

	}
}
