package proverb

import "fmt"

var lines = []string{
	"For want of a %v the %v was lost.",
	"For want of a %v the %v was lost.",
	"For want of a %v the %v was lost.",
	"For want of a %v the %v was lost.",
	"For want of a %v the %v was lost.",
	"For want of a %v the %v was lost.",
	"And all for the want of a %v.",
}

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return rhyme
	}
	out := []string{}
	for i := 0; i < len(lines); i++ {
		if i == len(rhyme)-1 {
			out = append(out, fmt.Sprintf(lines[len(lines)-1], rhyme[0]))
			return out
		}
		out = append(out, fmt.Sprintf(lines[i], rhyme[i], rhyme[i+1]))
	}
	return out
}
