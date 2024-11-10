package main

// Функция для сравнения двух строк после обработки символов backspace
func backspaceCompareV2(s string, t string) bool {
	/*
		METHOD: Two pointers
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	// Вспомогательная функция для нахождения следующего валидного индекса
	nextValidIndex := func(str string, index int) int {
		backspaceCount := 0
		for index >= 0 {
			if str[index] == '#' {
				backspaceCount++ // Увеличиваем счетчик backspace
			} else if backspaceCount > 0 {
				backspaceCount-- // Уменьшаем счетчик backspace, если есть символы для удаления
			} else {
				break // Выходим из цикла, если нет backspace для удаления
			}
			index--
		}
		return index
	}

	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		i = nextValidIndex(s, i)
		j = nextValidIndex(t, j)

		if i < 0 && j < 0 {
			return true // Если оба индекса меньше 0, строки эквивалентны
		}
		if i < 0 || j < 0 {
			return false // Если только один индекс меньше 0, строки не эквивалентны
		}
		if s[i] != t[j] {
			return false // Если символы не совпадают, строки не эквивалентны
		}

		i--
		j--
	}

	return true // Если все символы совпали, строки эквивалентны
}

// func main() {
// 	s := "ab#c"
// 	t := "ad#c"
// 	fmt.Println(backspaceCompare(s, t)) // Вывод: true

// 	s = "ab##"
// 	t = "c#d#"
// 	fmt.Println(backspaceCompare(s, t)) // Вывод: true

// 	s = "a##c"
// 	t = "#a#c"
// 	fmt.Println(backspaceCompare(s, t)) // Вывод: true

// 	s = "a#c"
// 	t = "b"
// 	fmt.Println(backspaceCompare(s, t)) // Вывод: false
// }
