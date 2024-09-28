package romannumerals

import (
	"errors"
	"math"
	"strconv"
	"unicode/utf8"
)

var numeralValues = map[int]rune{
	1:    'I',
	5:    'V',
	10:   'X',
	50:   'L',
	100:  'C',
	500:  'D',
	1000: 'M',
}

var arabicValues = []int{1000, 500, 100, 50, 10, 5, 1}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3_999 {
		return "", errors.New("make it > 0 and < 4000 next time")
	}

	numerals := []rune{}
	arabicString := strconv.Itoa(input)
	exponent := float64(utf8.RuneCountInString(arabicString) - 1)

	for _, r := range arabicString {
		switch r {
		case '0':
		case '4', '9':
			numerals = handleFourAndNine(numerals, r, exponent)
		default:
			numerals = handleDefault(numerals, r, exponent)
		}
		exponent--
	}

	return string(numerals), nil
}

func handleFourAndNine(numerals []rune, r rune, exponent float64) []rune {
	n, _ := strconv.Atoi(string(r))
	return append(
		numerals,
		numeralValues[int(math.Pow(10.0, exponent))],
		numeralValues[(n+1)*int(math.Pow(10.0, exponent))],
	)
}

func handleDefault(numerals []rune, r rune, exponent float64) []rune {
	n, _ := strconv.Atoi(string(r))
	n *= int(math.Pow(10.0, exponent))
	for _, v := range arabicValues {
		for divisions := n / v; divisions > 0; divisions-- {
			numerals = append(numerals, numeralValues[v])
			n -= v
		}
	}
	return numerals
}
