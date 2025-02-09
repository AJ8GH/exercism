package yacht

import "slices"

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return scoreNum(dice, 1)
	case "twos":
		return scoreNum(dice, 2)
	case "threes":
		return scoreNum(dice, 3)
	case "fours":
		return scoreNum(dice, 4)
	case "fives":
		return scoreNum(dice, 5)
	case "sixes":
		return scoreNum(dice, 6)
	case "full house":
		return fullHouse(dice)
	case "four of a kind":
		return ofKind(dice, 4)
	case "little straight":
		return straightWithMax(dice, 5)
	case "big straight":
		return straightWithMax(dice, 6)
	case "choice":
		return choice(dice)
	case "yacht":
		return ofKind(dice, 5)
	default:
		panic("bad strategy")
	}
}

func scoreNum(dice []int, num int) (score int) {
	for _, v := range dice {
		if v == num {
			score += v
		}
	}
	return score
}

func ofKind(dice []int, num int) (score int) {
	count := map[int]int{}
	for _, v := range dice {
		count[v]++
	}
	for k, v := range count {
		if v >= num {
			score += k * num
		}
		if num == 5 {
			score *= 2
		}
	}
	return score
}

func straightWithMax(dice []int, maxNum int) (score int) {
	switch {
	case choice(dice) == 15 && maxNum == 5 && slices.Contains(dice, 5):
		return 30
	case choice(dice) == 20 && maxNum == 6:
		return 30
	}
	return 0
}

func fullHouse(dice []int) (score int) {
	count := map[int]int{}
	for _, v := range dice {
		count[v]++
	}
	three := false
	two := false
	for _, v := range count {
		switch v {
		case 3:
			three = true
		case 2:
			two = true
		}
	}
	if three && two {
		return choice(dice)
	}
	return score
}

func choice(dice []int) (score int) {
	for _, v := range dice {
		score += v
	}
	return score
}
