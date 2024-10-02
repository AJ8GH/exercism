package lsproduct

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

var re = regexp.MustCompile(`^\d+$`)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	size := utf8.RuneCountInString(digits)
	if span > size || !re.Match([]byte(digits)) || span < 0 {
		return 0, errors.New("span bigger than size")
	}

	series := []string{}
	for i := 0; i <= size-span; i++ {
		series = append(series, digits[i:i+span])
	}
	return slices.Max(mapSlice(series, toScore)), nil
}

func toScore(series string) int64 {
	numStrings := strings.Split(series, "")
	var score int64 = 1
	for _, v := range numStrings {
		n, _ := strconv.Atoi(v)
		score *= int64(n)
	}
	return score
}

func mapSlice[T, U any](in []T, f func(T) U) (mapped []U) {
	mapped = []U{}
	for _, v := range in {
		mapped = append(mapped, f(v))
	}
	return
}
