package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	sum := 0
	nStr := strconv.Itoa(n)
	digits := float64(len(nStr))
	for _, v := range nStr {
		i, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(i), digits))
	}
	return sum == n
}
