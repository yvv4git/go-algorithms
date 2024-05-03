package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	/*
		METHOD: Two pointers

		TIME COMPLEXITY: O(n), где n - длина строк s и goal.
		Это обусловлено тем, что мы проходим по всем символам в строке s и goal не более двух раз.


		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество дополнительной памяти для хранения счётчика символов.
	*/
	// Проверяем, что длины строк s и goal одинаковы
	if len(s) != len(goal) {
		return false
	}

	// Проверяем, что строки s и goal совпадают
	if s == goal {
		// Проверяем, есть ли в строке s повторяющиеся символы
		count := make(map[rune]int)
		for _, char := range s {
			count[char]++
			if count[char] > 1 {
				return true
			}
		}
		return false
	}

	// Ищем два различных символа в s, которые должны быть такими же в goal, но в разном порядке
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	// Проверяем, что символы в s и goal на соответствующих позициях различаются ровно на один символ
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

func main() {
	fmt.Println(buddyStrings("ab", "ba"))               // true
	fmt.Println(buddyStrings("ab", "ab"))               // false
	fmt.Println(buddyStrings("aa", "aa"))               // true
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb")) // true
}
