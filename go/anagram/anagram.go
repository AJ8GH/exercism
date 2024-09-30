package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	out := []string{}

	for _, c := range candidates {
		s, lowerC := strings.ToLower(subject), strings.ToLower(c)
		if s == lowerC {
			continue
		}
		sRunes, cRunes := []rune(s), []rune(lowerC)
		slices.Sort(sRunes)
		slices.Sort(cRunes)
		if slices.Equal(sRunes, cRunes) {
			out = append(out, c)
		}
	}
	return out
}
