package _217_find_palindrome_with_fixed_length

import "math"

// KthPalindromeV2 - решение для задачи
// queries - порядковые номера, для которых надо сгенерировать палиндромы,
// intLength - какой длины должны быть палиндромы
func kthPalindromeV1(queries []int, intLength int) []int64 {
	/*
		METHOD: Math + DP
		Time complexity: O(n), это происходит потому, что функция проходит по каждому элементу массива queries и выполняет некоторые операции.
		Space complexity: O(n), это происходит потому, что функция создает новый массив result, размер которого равен размеру входного массива
	*/
	// Создаем слайс для хранения результатов.
	result := make([]int64, len(queries))

	// Вычисляем середину палиндрома.
	/*
		Формула start := int(math.Pow10(middle - 1)) используется для вычисления начала диапазона палиндромов.

		В общем случае, эта формула вычисляет начало диапазона палиндромов для заданной длины.
		Например, для длины 3 она вернет 100, для длины 4 - 1000, для длины 5 - 10000 и т.д.
	*/
	middle := (intLength + 1) / 2
	// Вычисляем начало диапазона палиндромов.
	start := int(math.Pow10(middle - 1))

	// Вычисляем конец диапазона палиндромов.
	var end int
	for i := 0; i < middle; i++ {
		end = end*10 + 9
	}

	// Генерируем палиндромы для каждого запроса.
	for i := 0; i < len(queries); i++ {
		// Если палиндром выходит за пределы допустимого диапазона, то записываем -1.
		if start+queries[i]-1 > end {
			result[i] = -1
			continue
		}

		// Генерируем палиндром и записываем его в результат.
		result[i] = genPalindrome(start+queries[i]-1, intLength%2 == 1)
	}

	// Возвращаем результат.
	return result
}

// Функция genPalindrome генерирует палиндром из заданной половины.
func genPalindrome(half int, odd bool) int64 {
	n := half
	if odd {
		n /= 10
	}

	// Создаем переменную для хранения результата.
	result := int64(half)
	// Добавляем вторую половину палиндрома.
	for ; n > 0; n /= 10 {
		result = result*10 + int64(n%10)
	}

	// Возвращаем результат.
	return result
}
