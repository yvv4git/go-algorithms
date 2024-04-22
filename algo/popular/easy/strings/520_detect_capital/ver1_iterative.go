package main

import (
	"fmt"
	"unicode"
)

/*
Функция detectCapitalUse должна определить, используется ли правильный регистр букв в слове. В зависимости от правильного использования регистра букв в слове, слово может быть написано как:
1. Все прописные буквы (например, "USA").
2. Все строчные буквы (например, "leetcode").
3. Первая буква прописная, остальные строчные (например, "Google").
*/
func detectCapitalUse(word string) bool {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(n), где n - количество букв в слове. Это связано с тем, что мы проходим по каждой букве в слове ровно один раз.

		SPACE COMPLEXITY: O(1), так как мы используем небольшое количество дополнительной памяти для хранения нескольких логических переменных, которые не зависят от размера входных данных.
	*/
	// Проверяем, является ли первая буква прописной
	firstCapital := unicode.IsUpper(rune(word[0]))

	// Проверяем, являются ли все остальные буквы прописными или строчными
	allCapital := true
	allLower := true
	for i := 1; i < len(word); i++ {
		if unicode.IsUpper(rune(word[i])) {
			allLower = false
		} else {
			allCapital = false
		}
	}

	return (firstCapital && allLower) || (allCapital && firstCapital) || allLower
}

func main() {
	//fmt.Println(detectCapitalUse("USA"))      // true
	//fmt.Println(detectCapitalUse("FlaG"))     // false
	//fmt.Println(detectCapitalUse("leetcode")) // true
	//fmt.Println(detectCapitalUse("Google"))   // true
	fmt.Println(detectCapitalUse("mL")) // false
}
