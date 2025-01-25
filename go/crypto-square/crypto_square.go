package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[a-zA-Z0-9]+`)

func Encode(pt string) string {
	letters := re.FindAllString(pt, -1)
	normalised := strings.ToLower(strings.Join(letters, ""))
	l := len(normalised)
	sq := int(math.Ceil(math.Sqrt(float64(l))))
	rows := rows(normalised, l, sq)

	var out []string
	for i := 0; i < sq; i++ {
		var chunk string
		for _, row := range rows {
			switch {
			case i < len(row):
				chunk += string(row[i])
			default:
				chunk += " "
			}
		}
		out = append(out, chunk)
	}

	return strings.Join(out, " ")
}

func rows(normalised string, l, sq int) (rows []string) {
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
	return rows
}
