package main

// Функция largestPalindromicV1 создает наибольший палиндром из цифр в строке num.
func largestPalindromicV1(num string) string {
	/*
		METHOD: Greedy algorithm
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)

		Time complexity
		Временная сложность данного алгоритма составляет O(n), где n - количество цифр в строке num.
		Это связано с тем, что мы проходим по всем цифрам в строке и выполняем некоторые операции для каждой цифры.

		Space complexity
		Пространственная сложность также составляет O(n), так как мы храним частоту каждой цифры в отдельном массиве.

		Жадный алгоритм - это алгоритм, который с каждым шагом делает локально оптимальный выбор, основываясь на информации, доступной в текущий момент.
		В данном случае, алгоритм пытается создать наибольший палиндром, добавляя цифры в порядке убывания их частоты.
	*/

	// Создаем массив частот для каждой цифры.
	freq := make([]int, 10)
	for _, ch := range num {
		freq[ch-'0']++
	}

	// Создаем префикс для палиндрома.
	prefix := ""
	// Создаем центральную цифру для палиндрома.
	centre := -1

	// Проходим по массиву частот в обратном порядке.
	// Идем в обратном порядке, чтобы получить максимальное число.
	// Таким образом, мы пытаемся максимизировать префикс палиндрома, используя все доступные пары цифр,
	// а также используя самую большую цифру, которая может быть использована для формирования палиндрома.
	for i := 9; i >= 0; i-- {
		// Пока частота текущей цифры больше 1, добавляем ее в префикс и уменьшаем частоту на 2.
		for freq[i] > 1 {
			prefix += string(rune('0' + i))
			freq[i] -= 2
		}
		// Если частота текущей цифры равна 1 и она больше центральной цифры, то обновляем центральную цифру.
		if freq[i] == 1 && i > centre {
			centre = i
		}
	}

	// Создаем центральную часть палиндрома.
	// Этот код нужен для того, чтобы добавить центральную цифру в середину палиндрома, если она существует.
	// Если центральной цифры нет, то палиндром будет симметричным относительно его середины.
	middle := ""
	if centre != -1 {
		middle = string(rune('0' + centre))
	}

	// Если префикс не пустой и его первый символ равен '0', то возвращаем центральную часть палиндрома.
	// В целом, этот код нужен для того, чтобы обработать случай, когда префикс - это "0", и центральная цифра была найдена.
	// В этом случае, мы возвращаем центральную цифру, которая может быть использована в середине палиндрома.
	if len(prefix) > 0 && prefix[0] == '0' {
		if centre == -1 {
			return "0"
		}
		return middle
	}

	// Возвращаем палиндром, состоящий из префикса, центральной части и префикса в обратном порядке.
	return prefix + middle + reverse(prefix)
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
