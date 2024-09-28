package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("make it > 0 and < 64 next time")
	}
	out := uint64(1)
	for i := 0; i < number-1; i++ {
		out *= 2
	}
	return out, nil
}

func Total() uint64 {
	sum := uint64(0)
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum += v
	}
	return sum
}
