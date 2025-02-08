package diamond

import "errors"

const (
	first = 'A'
	last  = 'Z'
)

func Gen(char byte) (string, error) {
	if char < first || char > last {
		return "", errors.New("bad char")
	}
	if char == first {
		return string(first), nil
	}

	diff := int(char) - int(first)
	padding := diff*2 - 1

	var out string
	for i := first; i <= rune(char); i++ {
		out += buildLine(char, i, diff, padding)
	}

	for i := rune(char - 1); i >= first; i-- {
		out += buildLine(char, i, diff, padding)
	}
	return out[:len(out)-1], nil
}

func buildLine(char byte, i rune, diff, padding int) string {
	var line string
	switch i {
	case first:
		line = buildFirstLine(diff)
	case rune(char):
		line += string(char)
		for j := 0; j < padding; j++ {
			line += " "
		}
		line += string(char)
	default:
		line = buildRegularLine(i, padding)
	}
	return line + "\n"
}

func buildRegularLine(currentChar rune, padding int) string {
	line := ""
	currentDiff := currentChar - first
	currentPadding := currentDiff*2 - 1
	middle := string(currentChar)
	for j := 0; j < int(currentPadding); j++ {
		middle += " "
	}
	middle += string(currentChar)
	sidePadding := padding - int(currentPadding)
	for j := 0; j < sidePadding/2; j++ {
		line += " "
	}
	line += middle
	for j := 0; j < sidePadding/2; j++ {
		line += " "
	}
	return line
}

func buildFirstLine(diff int) string {
	line := ""
	for j := 0; j < diff; j++ {
		line += " "
	}
	line += string(first)
	for j := 0; j < diff; j++ {
		line += " "
	}
	return line
}
