package rotationalcipher

import (
	"strings"
	"unicode"
)

const lettersInAlphabet = 26

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func RotationalCipher(plain string, shiftKey int) string {
	out := []rune{}
	for _, r := range plain {
		rotated := rotate(r, shiftKey)
		out = append(out, rotated)
	}
	return string(out)
}

func rotate(r rune, shiftKey int) rune {
	rotate := unicode.ToLower(r)
	if !strings.ContainsRune(alphabet, rotate) {
		return r
	}

	currentIndex := strings.IndexRune(alphabet, rotate)
	newIndex := currentIndex + shiftKey
	newIndex %= lettersInAlphabet
	rotate = rune(alphabet[newIndex])

	if unicode.IsUpper(r) {
		rotate = unicode.ToUpper(rotate)
	}

	return rotate
}
