package allyourbase

import (
	"errors"
	"math"
	"slices"
)

func ConvertToBase(
	inputBase int,
	inputDigits []int,
	outputBase int,
) (
	out []int,
	err error,
) {
	switch {
	case inputBase < 2:
		return nil, errors.New("input base must be >= 2")
	case outputBase < 2:
		return nil, errors.New("output base must be >= 2")
	case invalid(inputDigits, inputBase):
		return nil, errors.New("all digits must satisfy 0 <= d < input base")
	case len(inputDigits) == 0 || allZeroes(inputDigits):
		return []int{0}, nil
	}

	sum := 0.0
	i := 0.0
	for j := len(inputDigits) - 1; j >= 0; j-- {
		base := float64(inputBase)
		n := float64(inputDigits[j])
		factor := (math.Pow(base, i))
		sum += n * factor
		i++
	}

	sumInt := int(sum)
	for sumInt != 0 {
		digit := sumInt % outputBase
		out = append(out, digit)
		sumInt /= outputBase
	}

	slices.Reverse(out)
	return out, nil
}

func allZeroes(digits []int) bool {
	for _, v := range digits {
		if v > 0 {
			return false
		}
	}
	return true
}

func invalid(digits []int, inputBase int) bool {
	for _, d := range digits {
		if d < 0 || d >= inputBase {
			return true
		}
	}
	return false
}
