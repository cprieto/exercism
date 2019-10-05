package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	freqMap := FreqMap{}
	freqStream := make(chan FreqMap, len(s))

	for _, p := range s {
		go func(s string) {
			freqStream <- Frequency(s)
		}(p)
	}

	for range s {
		for k, v := range <- freqStream {
			freqMap[k] += v
		}
	}

	return freqMap
}
