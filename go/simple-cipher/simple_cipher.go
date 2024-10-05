package cipher

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode/utf8"
)

const lowerA = 97
const lowerZ = 122
const alphabetSize = 26

var reInput = regexp.MustCompile(`[a-z]+`)
var reKeyValid = regexp.MustCompile(`^[a-z]+$`)
var reKeyInvalid = regexp.MustCompile(`^[a]+$`)

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance >= alphabetSize || distance <= -alphabetSize {
		return nil
	}
	return shift{distance}
}

func (c shift) Encode(input string) string {
	return code(input, c.distance, add)
}

func (c shift) Decode(input string) string {
	return code(input, c.distance, subtract)
}

func code(input string, distance int, operator func(a, b int) int) string {
	encoded := []rune{}
	for _, v := range strings.ToLower(input) {
		shifted := shiftRune(v, distance, operator)
		if shifted != -1 {
			encoded = append(encoded, shifted)
		}
	}
	return string(encoded)
}

func shiftRune(r rune, distance int, operator func(a, b int) int) rune {
	if r >= lowerA && r <= lowerZ {
		shifted := operator(int(r), distance)
		if shifted < lowerA {
			shifted += alphabetSize
		} else if shifted > lowerZ {
			shifted -= alphabetSize
		}
		return rune(shifted)
	}
	return -1
}

func add(r, distance int) int {
	return r + distance
}

func subtract(r, distance int) int {
	return r - distance
}

func NewVigenere(key string) Cipher {
	if !reKeyValid.MatchString(key) || reKeyInvalid.MatchString(key) {
		return nil
	}
	if len(key) == 0 {
		key = generateKey()
	}
	return vigenere{key}
}

func generateKey() string {
	key := []rune{}
	for i := 0; i < 100; i++ {
		key = append(key, rune(rand.Intn(alphabetSize)+lowerA))
	}
	return string(key)
}

func (v vigenere) Encode(input string) string {
	return codeVigenere(input, v.key, add)
}

func (v vigenere) Decode(input string) string {
	return codeVigenere(input, v.key, subtract)
}

func codeVigenere(input, key string, operator func(a, b int) int) string {
	input = strings.Join(reInput.FindAllString(strings.ToLower(input), -1), "")
	encoded := []rune{}
	for i, r := range input {
		keyVal := key[i%utf8.RuneCountInString(key)]
		distance := int(keyVal - lowerA)
		shifted := shiftRune(r, distance, operator)
		if shifted != -1 {
			encoded = append(encoded, shifted)
		}
	}
	return string(encoded)
}
