package _177_can_make_palindrome_from_substring

// Функция для подсчета количества вхождений символов в подстроку
func countChars(s string, left, right int) [26]int {
	counts := [26]int{}
	for i := left; i <= right; i++ {
		counts[s[i]-'a']++
	}
	return counts
}

// Функция для проверки возможности создания палиндрома из подстроки
// Временная сложность: O(n), где n - длина подстроки
// Пространственная сложность: O(1), так как в худшем случае мы можем хранить 26 символов (букв английского алфавита)
func canMakePaliQueriesV1(s string, queries [][]int) []bool {
	/*
		METHOD: Prefix sums
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность решения составляет O(n), где n - длина подстроки, так как мы проходимся по подстроке дважды:
		один раз для подсчета количества вхождений символов, а второй раз для проверки возможности создания палиндрома.

		Space complexity
		Пространственная сложность составляет O(1), так как в худшем случае мы можем хранить 26 символов (букв английского алфавита).


		Для решения задачи используется подход "Префиксные суммы".
		Мы будем подсчитывать количество вхождений каждого символа в подстроку и использовать эти суммы для проверки возможности создания палиндрома.

		В этом решении мы используем функцию countChars для подсчета количества вхождений символов в подстроку.
		Затем мы проверяем, можно ли составить палиндром из подстроки, используя подсчитанные количества вхождений символов.
		Если количество символов, которые встречаются нечетное количество раз, меньше или равно k, то мы можем составить палиндром.
	*/
	// Вычисляем длину строки s
	n := len(s)

	// Создаем массив префиксных сумм размером n+1, где каждый элемент - это массив из 26 элементов типа int
	// Каждый элемент массива prefixSums[i] содержит количество вхождений каждой буквы алфавита до i-го символа в строке s
	prefixSums := make([][26]int, n+1)

	// Заполняем массив префиксных сумм
	for i := 0; i < n; i++ {
		// Копируем значения из предыдущего префикса
		prefixSums[i+1] = prefixSums[i]
		// Увеличиваем количество вхождений текущей буквы
		prefixSums[i+1][s[i]-'a']++
	}

	// Создаем массив ответов размером, равным количеству запросов
	result := make([]bool, len(queries))

	// Обрабатываем каждый запрос
	for i, query := range queries {
		// Извлекаем границы подстроки и количество операций замены
		left, right, k := query[0], query[1], query[2]

		// Создаем массив для хранения количества вхождений каждой буквы в подстроку
		counts := [26]int{}

		// Вычисляем количество вхождений каждой буквы в подстроку
		for j := 0; j < 26; j++ {
			counts[j] = prefixSums[right+1][j] - prefixSums[left][j]
		}

		// Вычисляем количество букв, которые встречаются нечетное количество раз
		oddCount := 0
		for j := 0; j < 26; j++ {
			if counts[j]%2 == 1 {
				oddCount++
			}
		}

		// Проверяем, можно ли составить палиндром из подстроки с помощью операций замены
		result[i] = k >= (oddCount / 2)
	}

	// Возвращаем массив ответов
	return result
}
