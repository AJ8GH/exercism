package acronym

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[a-zA-Z']+`)

func Abbreviate(s string) string {
	matches := re.FindAllString(s, -1)
	acronym := []rune{}
	for _, v := range matches {
		acronym = append(acronym, rune(v[0]))
	}
	return strings.ToUpper(string(acronym))
}
