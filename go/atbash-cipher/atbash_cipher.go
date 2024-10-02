package atbash

import (
	"regexp"
	"strings"
)

var plain = "abcdefghijklmnopqrstuvwxyz0123456789"
var cipher = "zyxwvutsrqponmlkjihgfedcba0123456789"
var re = regexp.MustCompile(`\w`)

func Atbash(s string) string {
	out := []string{}
	for i, v := range re.FindAllString(strings.ToLower(s), -1) {
		if i%5 == 0 && i > 0 {
			out = append(out, " ")
		}
		out = append(out, string(cipher[strings.Index(plain, v)]))
	}
	return strings.Join(out, "")
}
