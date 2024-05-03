package main

func buddyStringsV2(s string, goal string) bool {
	// Проверяем, что длины строк s и goal одинаковы
	if len(s) != len(goal) {
		return false
	}

	// Проверяем, что строки s и goal отличаются ровно на два места
	diff := make([]int, 0, 2)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}

	// Если строки одинаковые, проверяем наличие повторяющегося символа
	if len(diff) == 0 {
		count := make(map[rune]int)
		for _, char := range s {
			count[char]++
			if count[char] > 1 {
				return true
			}
		}
		return false
	}

	// Если строки отличаются ровно на два места, проверяем, что символы, которые отличаются, в обратном порядке должны быть в обеих строках
	if len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]] {
		return true
	}

	return false
}

//func main() {
//	fmt.Println(buddyStrings("ab", "ba"))               // true
//	fmt.Println(buddyStrings("ab", "ab"))               // false
//	fmt.Println(buddyStrings("aa", "aa"))               // true
//	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb")) // true
//}
