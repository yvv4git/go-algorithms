package _7_letter_combinations_of_a_phone_number

func letterCombinationsV3(digits string) []string {
	/*
		METHOD: DFS
		Time complexity: O(4^n), где n - количество цифр в входной строке.
		Space complexity: O(4^n), так как в худшем случае мы можем иметь 4^n комбинаций.

		Time complexity
		Временная сложность O(4^n) и пространственная сложность O(4^n) в этих решениях связаны с тем,
		что для каждой цифры в входной строке мы можем выбрать 3-4 буквы.
		Это означает, что для каждой цифры мы можем создать 3-4 новые комбинации, и так для каждой цифры в строке.
		Например, если у нас есть две цифры, то мы можем создать 34 = 12 новых комбинаций.
		Если у нас есть три цифры, то мы можем создать 34*4 = 48 новых комбинаций.

		Space complexity
		Пространственная сложность будет также O(4^n), так ка в худшем случае мы можем иметь 4^n комбинаций.
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

	var result []string
	dfs(digits, 0, "", &result, phone)

	return result
}

func dfs(digits string, index int, path string, result *[]string, phone map[string]string) {
	if index == len(digits) {
		*result = append(*result, path)
		return
	}

	for _, letter := range phone[string(digits[index])] {
		dfs(digits, index+1, path+string(letter), result, phone)
	}
}
