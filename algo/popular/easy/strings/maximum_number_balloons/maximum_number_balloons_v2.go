package maximum_number_balloons

const balon = "balon"

func maxNumberOfBalloonsV2(text string) int {
	/*
		METHOD: Hashing
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	mm := make(map[byte]int)
	for i := range text {
		mm[text[i]]++
	}
	mm['l'] /= 2
	mm['o'] /= 2

	count := 100000

	for i := range balon {
		if v := mm[balon[i]]; v < count {
			count = v
		}
	}

	return count
}
