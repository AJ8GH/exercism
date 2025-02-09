package foodchain

import (
	"fmt"
	"strings"
)

const (
	prefix           = "I know an old lady who swallowed a %s."
	toCatch          = "She swallowed the %s to catch the %s%s"
	spiderConnector1 = "It"
	spiderConnector2 = " that"

	flyPostfix    = "I don't know why she swallowed the fly. Perhaps she'll die."
	spiderPostfix = "%s wriggled and jiggled and tickled inside her."
	birdPostfix   = "How absurd to swallow a bird!"
	catPostfix    = "Imagine that, to swallow a cat!"
	dogPostfix    = "What a hog, to swallow a dog!"
	goatPostfix   = "Just opened her throat and swallowed a goat!"
	cowPostfix    = "I don't know how she swallowed a cow!"
	horsePostfix  = "She's dead, of course!"

	fly    = "fly"
	spider = "spider"
	bird   = "bird"
	cat    = "cat"
	dog    = "dog"
	goat   = "goat"
	cow    = "cow"
	horse  = "horse"
)

var postfixes = map[string]string{
	fly:    flyPostfix,
	spider: spiderPostfix,
	bird:   birdPostfix,
	cat:    catPostfix,
	dog:    dogPostfix,
	goat:   goatPostfix,
	cow:    cowPostfix,
	horse:  horsePostfix,
}

var animals = []string{
	fly,
	spider,
	bird,
	cat,
	dog,
	goat,
	cow,
	horse,
}

func Verse(v int) string {
	return buildVerse(v)
}

func Verses(start, end int) (out string) {
	for i := start; i <= end; i++ {
		out += Verse(i) + "\n\n"
	}
	return strings.TrimSpace(out)
}

func Song() string {
	return Verses(1, 8)
}

func buildVerse(v int) string {
	animal := animals[v-1]
	pre := fmt.Sprintf(prefix, animal)
	var post string
	switch animal {
	case spider:
		post = fmt.Sprintf(spiderPostfix, spiderConnector1)
	default:
		post = postfixes[animal]
	}

	if animal == horse {
		return strings.TrimSpace(fmt.Sprintf("%s\n%s", pre, post))
	}
	var acc string
	for i := v - 1; i > 0; i-- {
		acc += accumulator(i) + "\n"
	}

	return strings.TrimSpace(fmt.Sprintf("%s\n%s\n%s", pre, post, acc))
}

func accumulator(v int) string {
	animal1 := animals[v]
	animal2 := animals[v-1]

	postFix := "."
	switch animal2 {
	case spider:
		postFix = fmt.Sprintf(spiderPostfix, spiderConnector2)
	case fly:
		postFix += "\n" + flyPostfix
	}

	return fmt.Sprintf(toCatch, animal1, animal2, postFix)
}
