package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	/*
		METHOD: COUNTING SORT
		DESCRIPTION: Используем метод подсчета (Counting Sort) для определения количества элементов, меньших текущего.

		TIME COMPLEXITY: O(n + k), где n - количество элементов в массиве nums, а k - диапазон значений (в данном случае 101).
		SPACE COMPLEXITY: O(k), где k - диапазон значений (в данном случае 101).
	*/
	// Создаём массив для хранения количества вхождений каждого числа
	count := make([]int, 101)

	// Подсчитываем количество вхождений каждого числа в массиве nums
	for _, num := range nums {
		count[num]++
	}

	// Преобразуем массив count так, чтобы для каждого числа хранилось количество элементов, меньших этого числа
	for i := 1; i < 101; i++ {
		count[i] += count[i-1]
	}

	// Создаём результирующий массив для хранения ответа
	result := make([]int, len(nums))

	// Заполняем результирующий массив
	for i, num := range nums {
		if num > 0 {
			result[i] = count[num-1]
		} else {
			result[i] = 0
		}
	}

	return result
}

func main() {
	// Пример ввода
	nums := []int{8, 1, 2, 2, 3}

	// Вызываем функцию и получаем результат
	result := smallerNumbersThanCurrent(nums)

	// Выводим результат
	fmt.Println(result) // Output: [4, 0, 1, 1, 3]
}
