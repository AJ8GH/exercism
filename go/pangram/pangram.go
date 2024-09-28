package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	runes := map[rune]bool{}
	re := regexp.MustCompile(`[^A-Z]+`)
	for _, v := range re.ReplaceAllString(strings.ToUpper(input), "") {
		runes[v] = true
		if len(runes) == 26 {
			return true
		}
	}
	return false
}
