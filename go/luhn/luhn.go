package luhn

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

func Valid(id string) bool {
	re := regexp.MustCompile(`^[\d]+$`)
	id = regexp.MustCompile(`[\s\t]+`).ReplaceAllString(id, "")
	if !re.MatchString(id) || utf8.RuneCountInString(id) <= 1 {
		return false
	}

	lastRuneIndex := utf8.RuneCountInString(id) - 1

	doublingIndex := 0
	if lastRuneIndex%2 == 0 {
		doublingIndex = 1
	}

	sum := 0
	for i := lastRuneIndex; i >= 0; i-- {
		v, _ := strconv.Atoi(string(id[i]))
		if i%2 == doublingIndex {
			v = double(v)
		}
		sum += v
	}
	return sum%10 == 0
}

func double(n int) int {
	n *= 2
	if n > 9 {
		n -= 9
	}
	return n
}
