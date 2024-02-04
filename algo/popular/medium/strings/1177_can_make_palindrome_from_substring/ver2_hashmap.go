package _177_can_make_palindrome_from_substring

func canMakePaliQueriesV2(s string, queries [][]int) []bool {
	/*
		METHOD: Hashmap
		TIME COMPLEXITY: O(n + m), где n - длина строки s, m - количество запросов
		SPACE COMPLEXITY: O(n), где n - длина строки s.

		Объяснение:
		1. Создаем массив prefixCounts, где prefixCounts[i] - это количество вхождений каждого символа до индекса i в строке s.
		2. Для каждого запроса вычисляем количество нечетных символов в подстроке s[left:right+1].
		3. Проверяем, можно ли составить палиндром из подстроки s[left:right+1] с помощью операций замены символов.

		Преимущества:
		1. Этот метод эффективен, поскольку он использует префиксные суммы для подсчета количества вхождений каждого символа в подстроку.
		2. Он позволяет проверить, можно ли составить палиндром из подстроки с помощью операций замены символов за константное время.
	*/
	n := len(s)
	prefixCounts := make([][26]int, n+1)
	for i := 0; i < n; i++ {
		copy(prefixCounts[i+1][:], prefixCounts[i][:])
		prefixCounts[i+1][s[i]-'a']++
	}

	results := make([]bool, len(queries))
	for i, query := range queries {
		left, right, k := query[0], query[1], query[2]
		oddCounts := 0
		for j := 0; j < 26; j++ {
			count := prefixCounts[right+1][j] - prefixCounts[left][j]
			oddCounts += count % 2
		}
		results[i] = k*2 >= n-right+left+1 && k >= oddCounts/2
	}
	return results
}
