package main

func maxPowerV4(s string) int {
	counts := make(map[byte]int)
	currentChar := s[0]
	currentCount := 1
	maxCount := 1

	for i := 1; i < len(s); i++ {
		if s[i] == currentChar {
			currentCount++
		} else {
			counts[currentChar] = max(counts[currentChar], currentCount)
			currentChar = s[i]
			currentCount = 1
		}
	}
	counts[currentChar] = max(counts[currentChar], currentCount)

	for _, count := range counts {
		maxCount = max(maxCount, count)
	}
	return maxCount
}
