// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var ws = regexp.MustCompile(`^[\s\t\n\r]*$`)
var hl = regexp.MustCompile(`[a-zA-Z]+`)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
	case ws.MatchString(remark):
		return "Fine. Be that way!"
	case isQuestion(remark) && isYell(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYell(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isQuestion(s string) bool {
	s = strings.TrimSpace(s)
	return s[utf8.RuneCountInString(s)-1] == '?'
}

func isYell(s string) bool {
	return hl.MatchString(s) && s == strings.ToUpper(s)
}
