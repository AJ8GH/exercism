package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must not be zero")
	}
	for i := 0; true; i++ {
		if n == 1 {
			return i, nil
		}
		if n%2 == 0 {
			n /= 2
		} else {
			n *= 3
			n++
		}
	}
	return 0, errors.New("never reached one")
}
