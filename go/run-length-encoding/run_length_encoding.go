package encode

import (
	"regexp"
	"strconv"
)

func RunLengthEncode(input string) (encoded string) {
	if len(input) <= 1 {
		return input
	}

	current := rune(input[0])
	count := 1

	for i, v := range input {
		if i == 0 {
			continue
		}

		if v == current {
			count++
		}

		if i >= len(input)-1 || v != current {
			countStr := ""
			if count > 1 {
				countStr = strconv.Itoa(count)
			}
			encoded += countStr + string(current)
			if i == len(input)-1 && v != current {
				encoded += string(v)
			} else {
				count = 1
				current = v
			}
		}
	}

	return encoded
}

func RunLengthDecode(input string) (decoded string) {
	num := ""
	numRe := regexp.MustCompile(`\d`)
	for _, v := range input {
		if numRe.MatchString(string(v)) {
			num += string(v)
			continue
		} else {
			if num == "" {
				decoded += string(v)
			} else {
				n, _ := strconv.Atoi(num)
				for i := 0; i < n; i++ {
					decoded += string(v)
				}
				num = ""
			}
		}
	}
	return decoded
}
