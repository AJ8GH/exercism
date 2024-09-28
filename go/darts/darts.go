package darts

func Score(x, y float64) int {
	switch {
	case isInCircle(x, y, 1):
		return 10
	case isInCircle(x, y, 5):
		return 5
	case isInCircle(x, y, 10):
		return 1
	default:
		return 0
	}
}

func isInCircle(x, y float64, radius int) bool {
	return x*x+y*y <= float64(radius*radius)
}
