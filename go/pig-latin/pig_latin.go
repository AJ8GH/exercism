package piglatin

import (
	"regexp"
	"strings"
)

var (
	vowelRe = regexp.MustCompile(`^([aeiou]|xr|yt).*`)
	consRe  = regexp.MustCompile(`^([bcdfghjklmnpqrstvwxyz]+)([aeiou]+|yt)(.*)`)
	quRe    = regexp.MustCompile(`^(.*qu)(.*)`)
)

func Sentence(sentence string) string {
	translated := []string{}
	for _, word := range strings.Split(sentence, " ") {
		translated = append(translated, translate(word))
	}

	return strings.Join(translated, " ")
}

func translate(word string) string {
	switch {
	case vowelRe.MatchString(word):
		return word + "ay"
	case quRe.MatchString(word):
		matches := quRe.FindAllStringSubmatch(word, 2)
		return matches[0][2] + matches[0][1] + "ay"
	case consRe.MatchString(word):
		matches := consRe.FindAllStringSubmatch(word, 2)
		return matches[0][2] + matches[0][3] + matches[0][1] + "ay"
	default:
		first := string(word[0])
		return word[1:] + first + "ay"
	}
}
