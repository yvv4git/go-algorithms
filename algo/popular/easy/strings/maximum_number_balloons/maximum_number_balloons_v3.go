package maximum_number_balloons

import "math"

func maxNumberOfBalloonsV3(text string) int {
	/*
		Method: Hashing & Frequency
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	freqArr := [26]int{}

	for _, char := range text {
		freqArr[char-'a']++
	}

	arr := [5]int{
		freqArr['b'-'a'],
		freqArr['a'-'a'],
		freqArr['l'-'a'] / 2,
		freqArr['o'-'a'] / 2,
		freqArr['n'-'a'],
	}
	return findMin(arr)
}

func findMin(arr [5]int) int {
	min := math.MaxInt64
	for _, v := range arr {
		min = int(math.Min(float64(v), float64(min)))
	}
	return min
}
