package letter

import "sync"

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
	resFrequencies := FreqMap{}
	var wg sync.WaitGroup
	frequencies := make([]FreqMap, len(texts))
	for i := 0; i < len(texts); i++ {
		wg.Add(1)
		go func(l int) {
			frequencies[l] = Frequency(texts[l])
			wg.Done()
		}(i)
	}
	wg.Wait()
	for _, m := range frequencies {
		for k, v := range m {
			resFrequencies[k] += v
		}
	}

	return resFrequencies
}
