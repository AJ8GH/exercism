package summultiples

func SumMultiples(limit int, divisors ...int) int {
	mults := map[int]bool{}
	for _, v := range divisors {
		if v == 0 {
			continue
		}

		n := v
		for n < limit {
			mults[n] = true
			n += v
		}
	}
	sum := 0
	for k := range mults {
		sum += k
	}

	return sum
}
