package longest_palindrome

func longestPalindromeV3(s string) int {
	/*
		METHOD: Use hash map
		Time complexity: O(n)
		Space complexity: O(1) - используем только символы из алфавита, значит константа O(1), так как количество символов в алфавите ограничено
	*/
	myMap := make(map[rune]bool)
	max := 0
	for _, r := range s {
		if _, ok := myMap[r]; ok {
			max += 2
			delete(myMap, r)
		} else {
			myMap[r] = true
		}
	}

	if len(myMap) > 0 {
		max++
	}

	return max
}
