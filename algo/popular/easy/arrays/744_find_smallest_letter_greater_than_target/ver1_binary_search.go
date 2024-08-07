package main

import "fmt"

// Функция для поиска наименьшей буквы, которая строго больше заданной цели
func nextGreatestLetter(letters []byte, target byte) byte {
	/*
		METHOD: Binary search
		Мы ищем позицию, в которой должна быть наша буква, чтобы она была строго больше заданной.
		Бинарный поиск позволяет нам эффективно искать эту позицию, так как массив букв отсортирован.

		TIME COMPLEXITY: O(log N), где N - количество букв в массиве letters.
		Это обусловлено тем, что мы используем бинарный поиск, который каждый раз делит поиск пополам.

		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество переменных, не зависящих от размера входных данных.
	*/
	// Если целевая буква больше или равна последней букве в массиве,
	// то наименьшей буквай, которая строго больше цели, будет первая буква в массиве
	if target >= letters[len(letters)-1] {
		return letters[0]
	}

	// Инициализируем две переменные: left и right, которые будут указывать на границы поиска
	left, right := 0, len(letters)-1

	// Пока left не больше right, выполняем цикл
	for left < right {
		// Вычисляем середину текущего диапазона
		mid := left + (right-left)/2

		// Если буквы в середине меньше или равно цели, то ищем в правой половине
		if letters[mid] <= target {
			left = mid + 1
		} else {
			// Иначе ищем в левой половине
			right = mid
		}
	}

	// Возвращаем букву по индексу left, которая будет наименьшей буквой, которая строго больше цели
	return letters[left]
}

func main() {
	// Пример использования функции
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')
	result := nextGreatestLetter(letters, target)
	fmt.Printf("Буква, которая строго больше '%c': '%c'\n", target, result)
}
