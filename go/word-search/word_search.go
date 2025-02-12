package wordsearch

import "errors"

const (
	u  = "u"
	d  = "d"
	r  = "r"
	l  = "l"
	ur = "ur"
	ul = "ul"
	dr = "dr"
	dl = "dl"
)

var dirs = []string{
	u,
	d,
	r,
	l,
	ur,
	ul,
	dr,
	dl,
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	results := map[string][2][2]int{}
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			checkForWords(i, j, words, puzzle, results)
		}
	}

	for _, v := range words {
		_, exists := results[v]
		if !exists {
			return nil, errors.New("bad words")
		}
	}
	return results, nil
}

func checkForWords(
	i, j int,
	words []string,
	puzzle []string,
	results map[string][2][2]int,
) {
	for _, word := range words {
		_, exists := results[word]
		if exists {
			continue
		}
		checkForWord(i, j, word, puzzle, results)
	}
}

func checkForWord(
	y, x int,
	word string,
	puzzle []string,
	results map[string][2][2]int,
) {
	start := [2]int{x, y}
	xCopy := x
	yCopy := y
	for _, dir := range dirs {
		xCopy = x
		yCopy = y
		for i, r := range word {
			if yCopy < 0 || yCopy >= len(puzzle) || xCopy < 0 || xCopy >= len(puzzle[yCopy]) {
				break
			}
			if rune(puzzle[yCopy][xCopy]) != r {
				break
			}
			if i != len(word)-1 {
				yCopy, xCopy = next(yCopy, xCopy, dir)
			} else {
				results[word] = [2][2]int{start, {xCopy, yCopy}}
				return
			}
		}
	}
}

func next(i, j int, dir string) (int, int) {
	switch dir {
	case u:
		return i - 1, j
	case d:
		return i + 1, j
	case r:
		return i, j + 1
	case l:
		return i, j - 1
	case ur:
		return i - 1, j + 1
	case ul:
		return i - 1, j - 1
	case dr:
		return i + 1, j + 1
	case dl:
		return i + 1, j - 1
	default:
		return i, j
	}
}
