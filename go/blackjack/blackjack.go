package blackjack

const split = "P"
const hit = "H"
const stand = "S"
const win = "W"
const blackjack = 21

var cards = map[string]int{
	"ace":   11,
	"eight": 8,
	"two":   2,
	"nine":  9,
	"three": 3,
	"ten":   10,
	"four":  4,
	"jack":  10,
	"five":  5,
	"queen": 10,
	"six":   6,
	"king":  10,
	"seven": 7,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case card1 == "ace" && card2 == "ace":
		return split
	case sum(card1, card2) == blackjack && ParseCard(dealerCard) >= 10:
		return stand
	case sum(card1, card2) == blackjack:
		return win
	case isInRange(card1, card2, 17, 20):
		return stand
	case isInRange(card1, card2, 12, 16) && ParseCard(dealerCard) < 7:
		return stand
	default:
		return hit
	}
}

func sum(card1 string, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

func isInRange(card1 string, card2 string, min int, max int) bool {
	var val = sum(card1, card2)
	return val >= min && val <= max
}
