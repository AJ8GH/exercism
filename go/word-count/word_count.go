package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`[\w]+'?[\w]+|[\w]+`)

func WordCount(phrase string) Frequency {
	all := re.FindAllString(phrase, -1)
	count := map[string]int{}

	for _, word := range all {
		count[strings.ToLower(word)]++
	}
	return count
}
