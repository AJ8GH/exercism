package twelve

import (
	"fmt"
	"strings"
)

const firstVerse = "On the %v day of Christmas my true love gave to me: "

var lyrics = []string{
	"twelve Drummers Drumming, ",
	"eleven Pipers Piping, ",
	"ten Lords-a-Leaping, ",
	"nine Ladies Dancing, ",
	"eight Maids-a-Milking, ",
	"seven Swans-a-Swimming, ",
	"six Geese-a-Laying, ",
	"five Gold Rings, ",
	"four Calling Birds, ",
	"three French Hens, ",
	"two Turtle Doves, ",
	"and ",
	"a Partridge in a Pear Tree.",
}

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {
	out := []string{fmt.Sprintf(firstVerse, days[i])}
	start, end := len(lyrics)-i-1, len(lyrics)
	if i == 1 {
		start++
	}
	out = append(out, lyrics[start:end]...)
	return strings.Join(out, "")
}

func Song() string {
	out := []string{}
	for i := 1; i <= 12; i++ {
		out = append(out, Verse(i))
	}
	return strings.Join(out, "\n")
}
