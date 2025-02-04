package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	frequencies := FreqMap{}
	resultChan := make(chan FreqMap)

	for _, s := range texts {
		go func(s string) {
			resultChan <- Frequency(s)
		}(s)
	}

	for range texts {
		fm := <-resultChan
		for k, v := range fm {
			frequencies[k] += v
		}
	}
	return frequencies
}
