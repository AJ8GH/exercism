package sieve

var nonPrimes = map[int]bool{}

func Sieve(limit int) []int {
	primes := []int{}
	for i := 2; i <= limit; i++ {
		n := i
		if nonPrimes[n] {
			continue
		}

		primes = append(primes, n)
		sum := n
		for sum < limit {
			sum += n
			nonPrimes[sum] = true
		}
	}
	return primes
}
