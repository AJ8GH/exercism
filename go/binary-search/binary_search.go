package binarysearch

func SearchInts(list []int, key int) int {
	upper := len(list) - 1
	lower := 0
	i := len(list) / 2
	for upper >= lower {
		switch {
		case upper-lower < 0:
			return -1
		case upper-lower == 0 && list[i] != key:
			return -1
		case list[i] == key:
			return i
		case list[i] > key:
			upper = i - 1
			i = upper - (upper-lower)/2
		case list[i] < key:
			lower = i + 1
			i = lower + (upper-lower)/2
		}
	}
	return -1
}
