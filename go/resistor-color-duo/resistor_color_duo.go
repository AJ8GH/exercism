package resistorcolorduo

import "strconv"

var resistorColorMap = map[string]rune{
	"black":  '0',
	"brown":  '1',
	"red":    '2',
	"orange": '3',
	"yellow": '4',
	"green":  '5',
	"blue":   '6',
	"violet": '7',
	"grey":   '8',
	"white":  '9',
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	v, _ := strconv.Atoi(
		string([]rune{resistorColorMap[colors[0]], resistorColorMap[colors[1]]}),
	)
	return v
}
