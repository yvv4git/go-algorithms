package main

import (
	"math"
)

//// Функция для получения цифры по индексу
//func getDigit(num int, place int) int {
//	return (num / place) % 10
//}

// Функция для выполнения LSD Radix Sort
func lsdRadixSort(nums []int) {
	if len(nums) == 0 {
		return
	}

	// Находим максимальное число, чтобы определить количество разрядов
	maxNum := math.MinInt64
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// Выполняем LSD Radix Sort по каждому разряду
	place := 1
	for maxNum/place > 0 {
		countingSort(nums, place)
		place *= 10
	}
}

//// Функция для выполнения Counting Sort для конкретного разряда
//func countingSort(nums []int, place int) {
//	count := make([]int, 10) // Подсчет массив для десятичных чисел (0-9)
//	output := make([]int, len(nums))
//
//	// Подсчитываем количество вхождений каждого разряда
//	for _, num := range nums {
//		digit := getDigit(num, place)
//		count[digit]++
//	}
//
//	// Изменяем подсчет, чтобы подсчет содержал позиции этого элемента в выходном массиве
//	for i := 1; i < 10; i++ {
//		count[i] += count[i-1]
//	}
//
//	// Построение выходного массива
//	for i := len(nums) - 1; i >= 0; i-- {
//		num := nums[i]
//		digit := getDigit(num, place)
//		output[count[digit]-1] = num
//		count[digit]--
//	}
//
//	// Копируем отсортированный массив обратно в исходный
//	copy(nums, output)
//}

// Функция для нахождения максимального расстояния между элементами
func maximumGapV3(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// Сортируем массив с помощью LSD Radix Sort
	lsdRadixSort(nums)

	// Находим максимальное расстояние
	maxGap := 0
	for i := 1; i < len(nums); i++ {
		maxGap = max(maxGap, nums[i]-nums[i-1])
	}

	return maxGap
}

//// Функция для нахождения максимума из двух чисел
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	nums := []int{3, 6, 9, 1}
//	fmt.Println(maximumGap(nums)) // Вывод: 3
//}
