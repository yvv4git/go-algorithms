package maximum_number_balloons

func maxNumberOfBalloonsV1(text string) int {
	/*
		Method: Hashing
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	hashMap := map[byte]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(text); i++ {
		if _, ok := hashMap[text[i]]; ok {
			hashMap[text[i]]++
		}
	}

	result := len(text)
	for key, value := range hashMap {
		if key == 'l' || key == 'o' {
			result = min(result, value/2) // Делим на 2, так как в слове balloon эти символы встречаются 2 раза.
		} else {
			result = min(result, value)
		}
	}

	return result
}
