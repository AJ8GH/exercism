package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[a-zA-Z0-9]+`)

func Encode(pt string) string {
	/*

		l -> length of message
		c -> number of columns
		r -> number of rows,

		for rectangle  r  x  c  find the smallest possible integer  c  such that:

		r * c >= l
		c >= r
		c - r <= 1

		---

		"1, 2, 3 GO!"

		123go

		123
		go

		1g 2o 3

		---


		"12345678"

		123
		456
		78

		"147 258 36 "

	*/

	letters := re.FindAllString(pt, -1)
	normalised := strings.ToLower(strings.Join(letters, ""))
	l := len(normalised)
	sq := int(math.Ceil(math.Sqrt(float64(l))))

	var rows []string

	for i := 0; i < l; i += sq {
		end := i + sq
		if end > l {
			end = l
		}
		chunk := normalised[i:end]
		for i := 0; i < sq-len(chunk); i++ {
			chunk += " "
		}
		rows = append(rows, chunk)
	}

	out := ""
	for i := 0; i < sq; i++ {
		for _, row := range rows {
			if i < len(row) {
				out += string(row[i])
			}
		}
		if i == sq-1 {
			break
		}
		out += " "
	}

	return out
}
