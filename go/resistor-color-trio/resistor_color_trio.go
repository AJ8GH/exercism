package resistorcolortrio

import (
	"fmt"
	"math"
)

var values = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var ohms = map[int]string{
	0: "ohms",
	1: "kiloohms",
	2: "megaohms",
	3: "gigaohms",
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	out := values[colors[0]] * 10
	out += values[colors[1]]
	numZeroes := values[colors[2]]
	if colors[1] == "black" {
		numZeroes++
		out /= 10
	}
	ohms := ohms[numZeroes/3]
	out *= int(math.Pow10(numZeroes % 3))
	return fmt.Sprintf("%d %s", out, ohms)
}
