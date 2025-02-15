package transpose

import (
	"fmt"
)

func Transpose(input []string) (transposed []string) {
	maxLen := 0
	for _, row := range input {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	padded := []string{}
	padStr := fmt.Sprintf("-%ds", maxLen)
	for i, row := range input {
		paddedRow := row
		if i < len(input)-1 {
			paddedRow = fmt.Sprintf(fmt.Sprintf("%%%s", padStr), row)
		}
		padded = append(padded, paddedRow)
	}

	for i := 0; i < maxLen; i++ {
		transposedRow := ""
		for j, row := range padded {
			char := ""
			if j == 0 || !onlySpacesRemaining(padded, i) {
				if i < len(row) {
					char = string(row[i])
				}
			}
			transposedRow += char
		}
		transposed = append(transposed, transposedRow)
	}
	return transposed
}

func onlySpacesRemaining(padded []string, col int) bool {
	for i := 1; i < len(padded); i++ {
		for j := col; j < len(padded[i]); j++ {
			if string(padded[i][j]) != " " {
				return false
			}
		}
	}
	return true
}
