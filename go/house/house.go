package house

import (
	"fmt"
	"strings"
)

var things = []string{
	"%s the horse and the hound and the horn\n",
	"%s the farmer sowing his corn\n",
	"%s the rooster that crowed in the morn\n",
	"%s the priest all shaven and shorn\n",
	"%s the man all tattered and torn\n",
	"%s the maiden all forlorn\n",
	"%s the cow with the crumpled horn\n",
	"%s the dog\n",
	"%s the cat\n",
	"%s the rat\n",
	"%s the malt\n",
	"%s the house that Jack built.",
}

var verbs = []string{
	"that belonged to",
	"that kept",
	"that woke",
	"that married",
	"that kissed",
	"that milked",
	"that tossed",
	"that worried",
	"that killed",
	"that ate",
	"that lay in",
}

const firstPrefix = "This is"

func Verse(v int) (out string) {
	for i, thing := range things[len(things)-v:] {
		var pre string
		switch i {
		case 0:
			pre = firstPrefix
		default:
			pre = verbs[len(verbs)-v+i]
		}

		out += fmt.Sprintf(thing, pre)
	}
	return strings.TrimSpace(out)
}

func Song() (out string) {
	for i := 1; i <= len(things); i++ {
		out += Verse(i) + "\n\n"
	}
	return strings.TrimSpace(out)
}
