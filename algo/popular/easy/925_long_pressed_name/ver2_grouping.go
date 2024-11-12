package main

func isLongPressedNameV2(name string, typed string) bool {
	/*
		METHOD: Grouping Characters
		TIME COMPLEXITY: O(n + m), где n - длина строки name, m - длина строки typed
		SPACE COMPLEXITY: O(1)
	*/

	// Функция для группировки символов
	groupify := func(s string) [][]rune {
		if len(s) == 0 {
			return nil
		}
		var groups [][]rune
		currentGroup := []rune{rune(s[0])}
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				currentGroup = append(currentGroup, rune(s[i]))
			} else {
				groups = append(groups, currentGroup)
				currentGroup = []rune{rune(s[i])}
			}
		}
		groups = append(groups, currentGroup)
		return groups
	}

	nameGroups := groupify(name)
	typedGroups := groupify(typed)

	// Сравниваем группы
	if len(nameGroups) != len(typedGroups) {
		return false
	}
	for i := 0; i < len(nameGroups); i++ {
		if len(nameGroups[i]) > len(typedGroups[i]) || nameGroups[i][0] != typedGroups[i][0] {
			return false
		}
	}
	return true
}
