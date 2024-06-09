package main

import "fmt"

// Функция бинарного поиска в отсортированном сдвинутом массиве
func search(nums []int, target int) int {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(log n), так как в худшем случае мы делим массив пополам на каждом шаге, где n - количество элементов в массиве.
		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных, не зависящих от размера входных данных.
	*/
	left, right := 0, len(nums)-1 // Инициализируем два указателя: left и right

	for left <= right { // Пока left не больше right
		mid := left + (right-left)/2 // Вычисляем средний индекс

		if nums[mid] == target { // Если элемент по этому индексу равен искомому, возвращаем его индекс
			return mid
		}

		// Если левая половина отсортирована
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] { // Если искомый элемент в левой половине
				right = mid - 1 // Сдвигаем правую границу влево
			} else {
				left = mid + 1 // Иначе сдвигаем левую границу вправо
			}
		} else { // Если правая половина отсортирована
			if nums[mid] < target && target <= nums[right] { // Если искомый элемент в правой половине
				left = mid + 1 // Сдвигаем левую границу вправо
			} else {
				right = mid - 1 // Иначе сдвигаем правую границу влево
			}
		}
	}

	return -1 // Если элемент не найден, возвращаем -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target)) // Выводим индекс искомого элемента
}
