package prime

func Factors(n int64) []int64 {
	factors := []int64{}
	current := int(n)

	for i := 2; i <= current; i++ {
		if current%i == 0 {
			for current%i == 0 {
				current /= i
				factors = append(factors, int64(i))
			}
		}
	}

	return factors
}
