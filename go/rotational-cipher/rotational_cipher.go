package rotationalcipher

const lettersInAlphabet = 26
const lowerA = 97
const lowerZ = 122
const upperA = 65
const upperZ = 90

func RotationalCipher(plain string, shiftKey int) string {
	out := []rune{}
	for _, r := range plain {
		out = append(out, rotate(r, shiftKey))
	}
	return string(out)
}

func rotate(r rune, shiftKey int) (rotated rune) {
	if (r >= lowerA && r <= lowerZ) || (r >= upperA && r <= upperZ) {
		if (r <= upperZ && int(r)+shiftKey > upperZ) || (int(r)+shiftKey > lowerZ) {
			return rune(int(r) + shiftKey - lettersInAlphabet)
		}
		return rune(int(r) + shiftKey)
	}
	return r
}
