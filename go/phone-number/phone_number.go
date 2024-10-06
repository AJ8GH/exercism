package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const formatPattern = "(%v) %v-%v"

var re = regexp.MustCompile(
	`^\+?(\d?)[\s\-\.]*\(?([2-9]{1}\d{2})\)?[\s\-\.]*([2-9]{1}\d{2})[\s\-\.]*(\d{4})`,
)

func Number(phoneNumber string) (string, error) {
	matches, err := parseAndValidate(phoneNumber)
	if err != nil {
		return "", err
	}
	return strings.Join(matches[1:], ""), nil
}

func AreaCode(phoneNumber string) (string, error) {
	matches, err := parseAndValidate(phoneNumber)
	if err != nil {
		return "", err
	}
	return matches[1], nil
}

func Format(phoneNumber string) (string, error) {
	matches, err := parseAndValidate(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(formatPattern, matches[1], matches[2], matches[3]), nil
}

func parseAndValidate(phoneNumber string) ([]string, error) {
	if !re.MatchString(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	matches := re.FindAllStringSubmatch(phoneNumber, -1)[0][1:]
	countryCode := matches[0]
	if countryCode != "" && countryCode != "1" {
		return nil, errors.New("invalid country code")
	}

	total := 0
	for _, v := range matches {
		total += len(v)
	}

	if total == 11 && countryCode != "1" {
		return nil, errors.New("11 digit phone number must start with 1")
	}

	if total > 11 {
		return nil, errors.New("phone number has more than 11 digits")
	}
	return matches, nil
}
