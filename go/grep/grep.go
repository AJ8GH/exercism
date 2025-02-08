package grep

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var (
	number      = false
	file        = false
	insensitive = false
	invert      = false
	entire      = false
)

func Search(pattern string, flags, files []string) (results []string) {
	parseFlags(flags)
	defer resetFlags()
	for _, v := range files {
		b, _ := os.ReadFile(v)
		lines := strings.Split(string(b), "\n")
		for i, line := range lines {
			if line == "" {
				continue
			}
			results = appendIfMatch(i+1, len(files), v, line, pattern, results)
		}
	}
	return results
}

func parseFlags(flags []string) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			number = true
		case "-l":
			file = true
		case "-i":
			insensitive = true
		case "-v":
			invert = true
		case "-x":
			entire = true
		}
	}
}

func appendIfMatch(
	lineNum, numFiles int,
	fileName, line, pattern string,
	results []string,
) []string {
	matched := isMatch(line, pattern)
	if matched {
		results = appendResult(lineNum, numFiles, fileName, line, results)
	}
	return results
}

func isMatch(line, pattern string) bool {
	lineToCheck := line
	patternToCheck := pattern
	matched := false
	if insensitive {
		patternToCheck = strings.ToLower(patternToCheck)
		lineToCheck = strings.ToLower(lineToCheck)
	}
	if entire {
		matched = lineToCheck == patternToCheck
	} else {
		matched = strings.Contains(lineToCheck, patternToCheck)
	}
	if invert {
		return !matched
	}
	return matched
}

func appendResult(lineNum, numFiles int, fileName, line string, results []string) []string {
	if file {
		if slices.Contains(results, fileName) {
			return results
		}
		return append(results, fileName)
	}
	if number {
		line = fmt.Sprintf("%d:%s", lineNum, line)
	}
	if numFiles > 1 {
		line = fmt.Sprintf("%s:%s", fileName, line)
	}
	return append(results, line)
}

func resetFlags() {
	number = false
	file = false
	invert = false
	insensitive = false
	entire = false
}
