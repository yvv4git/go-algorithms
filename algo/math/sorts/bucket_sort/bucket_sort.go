package main

import (
	"fmt"
	"sort"
)

// Функция для вычисления индекса бакета для числа
func bucketIndex(num int, min int, max int, numBuckets int) int {
	return (num - min) / ((max - min + 1) / numBuckets)
}

// Bucket Sort
func bucketSort(nums []int) []int {
	/*
		METHOD: Bucket Sort

		TIME COMPLEXITY:
		- В лучшем случае, когда входные данные равномерно распределены по бакетам,
		сложность составляет O(n), где n - количество чисел для сортировки.

		- В худшем случае, когда все числа попадают в один бакет (или близко к одному), сложность может достигать O(n^2),
		если используется сортировка, которая не является стабильной, например, сортировка пузырьком.

		- В среднем случае, когда бакеты имеют примерно одинаковое количество элементов,
		временная сложность составляет O(n + k), где k - количество бакетов.

		SPACE COMPLEXITY:
		- В худшем случае, когда все числа попадают в один бакет, пространственная сложность может достигать O(n),
		так как для сортировки всех чисел в одном бакете может потребоваться дополнительное пространство.

		- В остальных случаях, пространственная сложность составляет O(n + k), где n - количество чисел, k - количество бакетов.
	*/
	if len(nums) == 0 {
		return nums
	}

	// Находим минимальное и максимальное значение в массиве
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Количество бакетов
	numBuckets := len(nums)
	// Создаем бакеты
	buckets := make([][]int, numBuckets)

	// Распределяем числа по бакетам
	for _, num := range nums {
		index := bucketIndex(num, min, max, numBuckets)
		buckets[index] = append(buckets[index], num)
	}

	// Сортируем числа в каждом бакете и собираем их обратно в исходный массив
	sorted := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket) // Используем встроенную сортировку Go для бакетов - IntroSort
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {
	nums := []int{329, 457, 657, 839, 436, 720, 355}
	sortedNums := bucketSort(nums)
	fmt.Println(sortedNums)
}
