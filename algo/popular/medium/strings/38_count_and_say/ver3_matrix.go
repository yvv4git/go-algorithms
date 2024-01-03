package _8_count_and_say

import "strconv"

// Функция countAndSay генерирует n-е число в последовательности "Count and Say"
func countAndSayV3(n int) string {
	/*
		Method: Matrix
		Time complexity: O(n*m), где n - количество чисел в последовательности, а m - средняя длина строки, полученной на предыдущем шаге.
		Space complexity: O(n*m), так как в худшем случае мы можем хранить все строки, которые мы получаем на каждом шаге.
	*/

	// Начинаем с числа 1
	result := "1"
	// Повторяем процесс n-1 раз
	for i := 1; i < n; i++ {
		// Генерируем следующее число на основе предыдущего
		result = countAndSayHelperV2(result)
	}
	return result
}

// Функция countAndSayHelper генерирует следующее число на основе предыдущего
func countAndSayHelperV2(s string) string {
	result := ""
	count := 1
	// Проходим по строке
	for i := 1; i < len(s); i++ {
		// Если текущий символ равен предыдущему, увеличиваем счетчик
		if s[i] == s[i-1] {
			count++
		} else {
			// Если символы разные, добавляем количество повторяющихся символов и сам символ в результат
			result += strconv.Itoa(count) + string(s[i-1])
			// Сбрасываем счетчик
			count = 1
		}
	}
	// Добавляем количество повторяющихся символов и сам символ для последнего символа в строке
	result += strconv.Itoa(count) + string(s[len(s)-1])
	return result
}
