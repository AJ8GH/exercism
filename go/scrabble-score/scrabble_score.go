package scrabble

import "strings"

func Score(word string) int {
	sum := 0
	for _, v := range strings.ToUpper(word) {
		sum += runeScore(v)
	}
	return sum
}

func runeScore(r rune) int {
	switch r {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
