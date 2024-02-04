package main

import "strconv"

func largestPalindromeV2(num string) string {
	/*
		METHOD: Greedy algorithm
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	frequency := make([]int, 10)
	for _, c := range num {
		frequency[c-'0']++
	}

	// Создаем префикс на основе информации о то, сколько какая цифра встретилась раз(частота)
	prefix := ""
	centre := -1 // Задаем -1, чтобы потом можно было отслеживать наличие середины в палиндроме
	for i := 9; i >= 0; i-- {
		// Пока частота текущей цифры больше 1, добавляем ее в префикс и уменьшаем частоту на 2.
		for frequency[i] > 1 {
			prefix += strconv.Itoa(i)
			frequency[i] -= 2
		}

		// Если частота текущей цифры равна 1 и она больше центральной цифры, то обновляем центральную цифру.
		if frequency[i] == 1 && i > centre {
			centre = i
		}
	}

	// Задаем значение цифры в середине
	middle := ""
	if centre != -1 {
		middle = strconv.Itoa(centre)
	}

	// Защита на случай, если в начале будет цифра 0
	if len(prefix) > 0 && prefix[0] == '0' {
		if centre == -1 {
			return "0"
		}
		return middle
	}

	// Создаем постфикс
	postfix := reverse(prefix)
	return prefix + middle + postfix
}
