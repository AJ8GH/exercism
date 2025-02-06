package allergies

var allergenMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var nums = []uint{128, 64, 32, 16, 8, 4, 2, 1}

func Allergies(allergies uint) (allergens []string) {
	if allergies > 256 {
		ref := 128
		for ref <= int(allergies) {
			ref *= 2
		}

		ref /= 2
		allergies %= uint(ref)
	}

	for _, v := range nums {
		if v <= allergies {
			allergies -= v
			allergens = append(allergens, allergenMap[v])
		}
	}

	return allergens
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, v := range Allergies(allergies) {
		if v == allergen {
			return true
		}
	}
	return false
}
