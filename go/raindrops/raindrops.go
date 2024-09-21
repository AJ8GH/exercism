package raindrops

import "strconv"

const pling = "Pling"
const plang = "Plang"
const plong = "Plong"

func Convert(number int) string {
	switch {
	case number%105 == 0:
		return pling + plang + plong
	case number%35 == 0:
		return plang + plong
	case number%21 == 0:
		return pling + plong
	case number%15 == 0:
		return pling + plang
	case number%7 == 0:
		return plong
	case number%5 == 0:
		return plang
	case number%3 == 0:
		return pling
	default:
		return strconv.Itoa(number)
	}
}
