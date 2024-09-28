package strand

var compliments = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	runes := []rune{}
	for _, v := range dna {
		runes = append(runes, compliments[v])
	}
	return string(runes)
}
