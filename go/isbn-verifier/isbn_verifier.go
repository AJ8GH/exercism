package isbn

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

var isValidFormat = regexp.MustCompile(`^[\d\-]*X?$`)
var toRemove = regexp.MustCompile(`[\-]*`)

func IsValidISBN(isbn string) bool {
	if !isValidFormat.MatchString(isbn) {
		return false
	}

	isbn = toRemove.ReplaceAllString(isbn, "")
	if utf8.RuneCountInString(isbn) != 10 {
		return false
	}

	sum := 0
	for i, v := range isbn {
		if v == 'X' {
			sum += 10
		} else {
			n, _ := strconv.Atoi(string(v))
			sum += (n * (10 - i))
		}
	}
	return sum%11 == 0
}
