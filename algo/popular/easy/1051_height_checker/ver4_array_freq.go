//go:build ignore

package main

import (
	"fmt"
)

func heightChecker(heights []int) int {
	/*
		METHOD: Frequency Array
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(k)

		Подробное объяснение:
		- МЕТОД: Frequency Array
		  Мы подсчитываем частоту встречаемости каждого элемента в исходном массиве. Затем, используя этот массив частот, мы можем определить правильный порядок элементов и сравнить его с исходным массивом.

		- ВРЕМЕННАЯ СЛОЖНОСТЬ: O(n)
		  Подсчет частоты встречаемости каждого элемента занимает O(n) времени. Проход по массиву для определения количества несовпадающих элементов также занимает O(n) времени. Таким образом, общая временная сложность составляет O(n).

		- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ: O(k)
		  Мы используем дополнительный массив размером k для подсчета частоты встречаемости каждого элемента. Таким образом, общая пространственная сложность составляет O(k).
	*/

	// Предположим, что диапазон значений роста студентов от 1 до 100
	k := 101
	count := make([]int, k)

	// Подсчитываем частоту встречаемости каждого элемента
	for _, height := range heights {
		count[height]++
	}

	// Инициализируем счетчик несовпадающих элементов
	countMismatches := 0

	// Проходим по массиву и сравниваем элементы с отсортированным порядком
	index := 0
	for height := 1; height < k; height++ {
		for count[height] > 0 {
			if heights[index] != height {
				countMismatches++
			}
			index++
			count[height]--
		}
	}

	// Возвращаем количество несовпадающих элементов
	return countMismatches
}

func main() {
	// Пример использования функции
	heights := []int{1, 1, 4, 2, 1, 3}
	result := heightChecker(heights)
	fmt.Println(result) // Вывод: 3
}
