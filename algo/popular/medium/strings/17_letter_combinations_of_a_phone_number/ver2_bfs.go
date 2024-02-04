package _7_letter_combinations_of_a_phone_number

func letterCombinationsV2(digits string) []string {
	/*
		METHOD: BFS
		TIME COMPLEXITY: O(4^n), где n - количество цифр в входной строке. Это связано с тем, что для каждой цифры мы можем выбрать 3-4 буквы.
		SPACE COMPLEXITY: O(4^n), так как в худшем случае мы можем иметь 4^n комбинаций.
	*/
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	result := []string{""}

	for _, digit := range digits {
		letters := phone[string(digit)]
		temp := []string{}

		for _, letter := range letters {
			for _, combination := range result {
				temp = append(temp, combination+string(letter))
			}
		}

		result = temp
	}

	return result
}
