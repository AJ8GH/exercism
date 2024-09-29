package protein

import (
	"errors"
	"unicode/utf8"
)

var ErrStop = errors.New("stop error")
var ErrInvalidBase = errors.New("invalid base error")

func FromRNA(rna string) ([]string, error) {
	out := []string{}
	for i := 0; i < utf8.RuneCountInString(rna); i += 3 {
		codon := string(rna[i : i+3])
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return out, nil
		}
		if err == ErrInvalidBase {
			return nil, err
		}
		out = append(out, protein)
	}
	return out, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
