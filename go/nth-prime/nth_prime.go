package prime

import "errors"

const firstPrime = 2

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("no primes less than 1 fool")
	}

	for i, primesFound := 1, 0; primesFound < n; i++ {
		if isPrime(i) {
			primesFound++
		}

		if primesFound == n {
			return i, nil
		}
	}

	return 0, errors.New("did not find the primes")
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= n/2+1; i++ {
		if n%i == 0 && n != i {
			return false
		}
	}
	return true
}
