package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	m := map[rune]bool{}
	re := regexp.MustCompile(`[-\s\t\n\r]+`)
	for _, v := range strings.ToLower(word) {
		if m[v] && !re.MatchString(string(v)) {
			return false
		}
		m[v] = true
	}
	return true
}
