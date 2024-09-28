package reverse

import (
	"strings"
)

func Reverse(input string) string {
	out := []string{}
	s := strings.Split(input, "")
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return strings.Join(out, "")
}
