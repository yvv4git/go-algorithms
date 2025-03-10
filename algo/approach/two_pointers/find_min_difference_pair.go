//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Функция для поиска пары чисел с минимальной разницей
func findMinDifferencePair(arr []int) (int, int) {
	/*
		METHOD: Two Pointers Approach
		1. Сортируем массив для того, чтобы использовать два указателя.
		2. Один указатель устанавливаем в начале массива, а второй — в конце.
		3. Находим минимальную разницу между элементами на указателях.
		4. Если разница между текущими элементами меньше минимальной, обновляем результат.
		5. Сдвигаем указатели для нахождения следующей минимальной разницы.

		TIME COMPLEXITY: O(nlogn) из-за сортировки массива
		SPACE COMPLEXITY: O(1) — дополнительные данные не используются.
	*/

	// Сортируем массив
	sort.Ints(arr)

	// Устанавливаем два указателя: один на начало, другой на конец массива
	left, right := 0, len(arr)-1

	// Инициализируем переменные для минимальной разницы и пары
	minDiff := int(^uint(0) >> 1) // максимально возможное значение
	var resultPair [2]int

	// Проходим по массиву
	for left < right {
		diff := arr[right] - arr[left]

		// Если разница меньше текущей минимальной, обновляем
		if diff < minDiff {
			minDiff = diff
			resultPair[0], resultPair[1] = arr[left], arr[right]
		}

		// Сдвигаем указатели в зависимости от текущей разницы
		if arr[right]-arr[left] > minDiff {
			right--
		} else {
			left++
		}
	}

	// Возвращаем пару с минимальной разницей
	return resultPair[0], resultPair[1]
}

/*
	Задача: Найти пару чисел с минимальной разницей в отсортированном массиве.

	Условия задачи:
	Дан отсортированный массив целых чисел. Требуется найти пару чисел из массива, которые имеют минимальную разницу.

	Пример:
	Вход:
		arr = [1, 3, 8, 10, 15]

	Выход:
		(8, 10)

	Алгоритм использует метод двух указателей для нахождения пары чисел с минимальной разницей.

	ТАЙМ-КОМПЛЕКСНОСТЬ: O(nlogn) из-за сортировки массива
	СПЕЙС-КОМПЛЕКСНОСТЬ: O(1) — дополнительные массивы не используются
*/

func main() {
	// Задаем отсортированный массив
	arr := []int{1, 3, 8, 10, 15}

	// Вызываем функцию для поиска пары с минимальной разницей
	num1, num2 := findMinDifferencePair(arr)

	// Выводим результат
	fmt.Printf("Пара с минимальной разницей: %d и %d\n", num1, num2)
}
