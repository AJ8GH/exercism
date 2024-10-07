package bottlesong

import (
	"fmt"
	"strings"
)

var lyrics = []string{
	"%v green bottle%v hanging on the wall,",
	"%v green bottle%v hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be %v green bottle%v hanging on the wall.",
}

var nums = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	out := []string{}
	for i := startBottles; i > startBottles-takeDown; i-- {
		n := nums[i]
		nn := nums[i-1]
		first, second := plurals(i)
		out = append(out, fmt.Sprintf(lyrics[0], strings.Title(n), first))
		out = append(out, fmt.Sprintf(lyrics[1], strings.Title(n), first))
		out = append(out, lyrics[2])
		out = append(out, fmt.Sprintf(lyrics[3], nn, second))
		out = append(out, "")
	}
	return out[:len(out)-1]
}

func plurals(numBottles int) (first, second string) {
	switch numBottles {
	case 2:
		return "s", ""
	case 1:
		return "", "s"
	default:
		return "s", "s"
	}
}
