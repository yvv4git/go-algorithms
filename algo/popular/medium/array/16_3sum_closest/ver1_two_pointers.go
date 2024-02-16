package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n^2), где n - количество элементов в массиве, так как мы перебираем каждый элемент массива и для каждого элемента используем двух указателей для поиска комбинаций.
		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных для хранения промежуточных результатов, не зависящих от размера входных данных.
	*/
	// Сортируем массив
	sort.Ints(nums)

	// Инициализируем переменные для ответа и минимальной разницы
	closestSum := nums[0] + nums[1] + nums[2]
	minDiff := abs(closestSum - target)

	// Итерируемся по массиву
	for i := 0; i < len(nums)-2; i++ {
		// Инициализируем два указателя - один на следующий элемент, второй на конец массива
		left, right := i+1, len(nums)-1

		// Итерируемся до тех пор, пока левый указатель не станет больше правого
		for left < right {
			// Вычисляем сумму текущей тройки чисел
			sum := nums[i] + nums[left] + nums[right]

			// Если сумма равна целевому значению, возвращаем ее
			if sum == target {
				return sum
			}

			// Вычисляем абсолютную разницу между суммой и целевым значением
			diff := abs(sum - target)

			// Если разница меньше минимальной, обновляем ближайшую сумму и минимальную разницу
			if diff < minDiff {
				minDiff = diff
				closestSum = sum
			}

			// Если сумма меньше целевого значения, двигаем левый указатель вправо
			if sum < target {
				left++
			} else {
				// Иначе двигаем правый указатель влево
				right--
			}
		}
	}

	// Возвращаем ближайшую сумму
	return closestSum
}

// Функция для вычисления абсолютной разницы между двумя числами
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target)) // Выводит: 2
}
