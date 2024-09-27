package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("lengths don't match")
	}

	count := 0
	runesA := []rune(a)
	runesB := []rune(b)
	for i := 0; i < utf8.RuneCountInString(a); i++ {
		if runesA[i] != runesB[i] {
			count++
		}
	}
	return count, nil
}
