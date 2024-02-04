package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(n), где n - длина входной строки s. Это обусловлено тем, что мы проходим по строке только один раз, и для каждого символа выполняем постоянное количество операций.
		Space complexity: O(n), так как мы используем дополнительный массив dp размером n+1 для хранения промежуточных результатов.

		Таким образом, временная и пространственная сложности этого алгоритма линейные, что является оптимальным решением для этой задачи.
	*/
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	// Создаем массив dp, где dp[i] будет содержать количество возможных декодированных сообщений для подстроки s[0:i]
	dp := make([]int, len(s)+1)
	dp[0] = 1 // пустая строка может быть декодирована в одно пустое сообщение

	if s[0] != '0' {
		dp[1] = 1 // если первый символ не равен '0', то он может быть декодирован в одну букву
	} else {
		dp[1] = 0 // если первый символ равен '0', то он не может быть декодирован
	}

	for i := 2; i <= len(s); i++ {
		// Проверяем, может ли текущая цифра (s[i-1]) быть декодирована как отдельная буква
		oneDigit, _ := strconv.Atoi(s[i-1 : i])
		if oneDigit >= 1 && oneDigit <= 9 {
			dp[i] += dp[i-1]
		}

		// Проверяем, может ли текущая цифра (s[i-2:i]) быть декодирована как одна буква
		twoDigits, _ := strconv.Atoi(s[i-2 : i])
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func main() {
	fmt.Println(numDecodings("12"))  // Вывод: 2
	fmt.Println(numDecodings("226")) // Вывод: 3
}
