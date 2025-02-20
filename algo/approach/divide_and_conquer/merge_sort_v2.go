//go:build ignore

package main

import "fmt"

// Функция для слияния двух отсортированных подмассивов в один отсортированный
func merge(arr []int, left, mid, right int) {
	/*
		Задача: Слить два отсортированных подмассива в один отсортированный массив.

		Подход:
		1. Создаем два временных массива для левой и правой части.
		2. Используем два указателя для слияния этих массивов.
		3. Сравниваем элементы и добавляем их в итоговый массив.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n), где n — количество элементов.
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(n)
	*/
	leftSize := mid - left + 1
	rightSize := right - mid

	// Создаем временные массивы
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	// Копируем данные в временные массивы
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	// Используем два указателя для слияния
	i, j, k := 0, 0, left
	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Копируем оставшиеся элементы в итоговый массив
	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}
	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// Рекурсивная функция для сортировки массива с использованием алгоритма Merge Sort
func mergeSort(arr []int, left, right int) {
	/*
		Задача: Реализовать сортировку слиянием.

		Подход:
		1. Разделить массив на две части.
		2. Рекурсивно отсортировать каждую часть.
		3. Слить отсортированные части в один отсортированный массив.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n log n)
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(n)
	*/
	if left < right {
		mid := (left + right) / 2

		// Рекурсивная сортировка левой и правой части
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		// Слияние отсортированных частей
		merge(arr, left, mid, right)
	}
}

func main() {
	// Пример: массив для сортировки
	arr := []int{38, 27, 43, 3, 9, 82, 10}

	// Выводим исходный массив
	fmt.Println("Исходный массив:", arr)

	// Вызываем MergeSort для сортировки массива
	mergeSort(arr, 0, len(arr)-1)

	// Выводим отсортированный массив
	fmt.Println("Отсортированный массив:", arr)
}
