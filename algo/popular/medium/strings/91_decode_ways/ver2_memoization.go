package main

import (
	"strconv"
)

var memo map[int]int

func numDecodingsV2(s string) int {
	/*
		METHOD: Memoization
		TIME COMPLEXITY: O(n), где n - длина входной строки s.
		Это обусловлено тем, что мы проходим по строке только один раз, и для каждого символа выполняем постоянное количество операций.
		Space complexity: O(n) в худшем случае, когда мы не можем использовать никакие из ранее вычисленных значений
		и заполняем карту memo для каждого символа в строке.
		В худшем случае, когда каждый символ уникален, размер карты memo будет равен длине строки s.
		Однако, в среднем и лучшем случаях карта будет заполняться меньше, так как мы будем использовать уже вычисленные значения.
	*/
	memo = make(map[int]int)
	return decodeWays(s, 0)
}

func decodeWays(s string, index int) int {
	// Если мы достигли конца строки, это означает, что мы нашли допустимое декодирование
	if index == len(s) {
		return 1
	}

	// Если текущий символ равен '0', это не допустимое декодирование
	if s[index] == '0' {
		return 0
	}

	// Если мы уже вычислили это значение, возвращаем его из кэша
	if val, ok := memo[index]; ok {
		return val
	}

	// Инициализируем количество возможных декодирований для текущего индекса
	ans := decodeWays(s, index+1)

	// Проверяем, может ли текущая цифра (s[index:index+2]) быть декодирована как одна буква
	if index+2 <= len(s) {
		twoDigits, _ := strconv.Atoi(s[index : index+2])
		if twoDigits >= 10 && twoDigits <= 26 {
			ans += decodeWays(s, index+2)
		}
	}

	// Сохраняем результат в кэше и возвращаем его
	memo[index] = ans
	return ans
}

//func main() {
//	fmt.Println(numDecodings("12"))  // Вывод: 2
//	fmt.Println(numDecodings("226")) // Вывод: 3
//}
