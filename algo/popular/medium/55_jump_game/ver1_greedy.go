package main

import "fmt"

// Функция canJump решает задачу "Jump Game".
func canJump(nums []int) bool {
	/*
		METHOD: Gready algorithm
		Подход, который используется в данном решении, называется "жадным алгоритмом" (greedy algorithm).
		Жадный алгоритм делает локально оптимальный выбор на каждом шаге с надеждой, что это приведет к глобальному оптимальному решению.
		В данном случае на каждом шаге мы выбираем максимально возможный индекс, который можно достичь, и проверяем, не достигли ли мы конца массива.

		Мы проходим по массиву и на каждом шаге обновляем максимальный индекс, который мы можем достичь на текущий момент.
		Если в какой-то момент текущий индекс превышает максимальный индекс, который мы можем достичь, то возвращаем false.
		Если же максимальный индекс, который мы можем достичь, становится больше или равен последнему индексу массива, то возвращаем

		TIME COMPLEXITY: O(n), где n — длина массива nums. Мы проходим по массиву один раз.

		SPACE COMPLEXITY: O(1), так как мы используем только константное количество дополнительной памяти.
	*/

	// Длина массива nums.
	n := len(nums)
	// Если массив состоит из одного элемента, то мы уже находимся в конце.
	if n == 1 {
		return true
	}

	// Максимальный индекс, который мы можем достичь на текущий момент.
	maxReach := 0

	// Проходим по всем элементам массива.
	for i := 0; i < n; i++ {
		// Если текущий индекс больше максимального, который мы можем достичь,
		// то возвращаем false, так как мы не можем продолжить.
		if i > maxReach {
			return false
		}

		// Обновляем максимальный индекс, который мы можем достичь.
		maxReach = max(maxReach, i+nums[i])

		// Если максимальный индекс, который мы можем достичь, больше или равен последнему индексу,
		// то возвращаем true, так как мы можем достичь конца массива.
		if maxReach >= n-1 {
			return true
		}
	}

	// Если мы прошли весь массив и не достигли конца, возвращаем false.
	return false
}

// Вспомогательная функция для нахождения максимума двух чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Примеры использования функции canJump.
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
}
