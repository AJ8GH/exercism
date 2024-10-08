package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\].*$)`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[\*-=~]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`.*".*(?i)(password).*".*`)
	sum := 0
	for _, v := range lines {
		if re.MatchString(v) {
			sum++
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[\s\t]+(\w+)`)
	for i, v := range lines {
		if re.MatchString(v) {
			matches := re.FindStringSubmatch(v)
			lines[i] = fmt.Sprintf("[USR] %v %v", matches[1], v)
		}
	}
	return lines
}
