package wordy

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`\-?\d+|multiplied|plus|minus|divided|cubed`)
var isDigit = regexp.MustCompile(`^\-?\d+$`)
var isOperator = regexp.MustCompile(`^multiplied|plus|minus|divided$`)

var binaryOperators = map[string]func(a, b int) int{
	"plus":       func(a, b int) int { return a + b },
	"minus":      func(a, b int) int { return a - b },
	"multiplied": func(a, b int) int { return a * b },
	"divided":    func(a, b int) int { return a / b },
}

func Answer(question string) (int, bool) {
	matches := re.FindAllString(question, -1)
	if len(matches) == 0 {
		return 0, false
	}
	var ints []int
	var bo func(a, b int) int
	for _, v := range matches {
		switch {
		case isDigit.MatchString(v):
			if len(ints) > 0 && bo == nil {
				return 0, false
			}
			n, _ := strconv.Atoi(v)
			ints = append(ints, n)
			if len(ints) == 2 {
				ints = []int{bo(ints[0], ints[1])}
				bo = nil
			}
		case isOperator.MatchString(v):
			if len(ints) == 0 || bo != nil {
				return 0, false
			}
			bo = binaryOperators[v]
		default:
			return 0, false
		}
	}

	if bo != nil {
		return 0, false
	}
	return ints[0], true
}
