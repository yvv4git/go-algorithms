package main

import "fmt"

// Функция repeatedNTimes принимает массив nums и возвращает элемент, который повторяется N раз.
func repeatedNTimes(nums []int) int {
	/*
		Method: Hash / Map / Dict
		Time: O(n)
		Space: O(n)
	*/
	// Создаем пустое множество для отслеживания уже встреченных элементов
	seen := make(map[int]bool)

	// Проходим по всем элементам массива
	for _, num := range nums {
		// Если элемент уже есть в множестве, значит это искомый элемент
		if seen[num] {
			return num
		}
		// Добавляем элемент в множество
		seen[num] = true
	}

	// Если по какой-то причине элемент не найден (хотя по условию задачи это невозможно), возвращаем -1
	return -1
}

func main() {
	// Пример использования функции
	nums := []int{1, 2, 3, 3}
	result := repeatedNTimes(nums)
	fmt.Println("Повторяющийся элемент:", result) // Вывод: Повторяющийся элемент: 3
}
