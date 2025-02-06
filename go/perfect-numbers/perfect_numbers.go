package perfect

import "errors"

type (
	Classification string
)

var ErrOnlyPositive = errors.New("n must be positive")

const (
	ClassificationDeficient = "Defi"
	ClassificationPerfect   = "Perf"
	ClassificationAbundant  = "Abun"
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	factors := []int64{1}
	for i := int64(2); i <= n/int64(2); i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	sum := int64(0)

	for _, v := range factors {
		sum += v
	}

	switch {
	case n == 1, sum < n:
		return ClassificationDeficient, nil
	case sum == n:
		return ClassificationPerfect, nil
	default:
		return ClassificationAbundant, nil
	}
}
