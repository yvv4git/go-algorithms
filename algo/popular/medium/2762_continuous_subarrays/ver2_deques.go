//go:build ignore

package main

import (
	"container/list"
	"fmt"
)

func continuousSubarrays(nums []int) int64 {
	/*
		METHOD: Deques
		- Используется подход с двумя деками: `maxDeque` и `minDeque`.
		- Дек `maxDeque` хранит индексы максимальных элементов в текущем окне.
		- Дек `minDeque` хранит индексы минимальных элементов в текущем окне.

		TIME COMPLEXITY: O(n)
		- Дек `maxDeque` и `minDeque` проходят через массив один раз, что занимает O(n).
		- Обновление деков и проверка условия выполняется за O(1) для каждого элемента.

		SPACE COMPLEXITY: O(n)
		- Используется фиксированное количество дополнительной памяти: переменные `count`, `left`, `maxDeque`, `minDeque`.
		- Никакие дополнительные структуры данных (например, массивы или хэш-таблицы) не используются.
	*/
	n := len(nums)
	if n == 0 {
		return 0
	}

	var count int64 = 0
	left := 0

	// Дек для хранения индексов максимальных элементов
	maxDeque := list.New()
	// Дек для хранения индексов минимальных элементов
	minDeque := list.New()

	for right := 0; right < n; right++ {
		// Обновляем maxDeque
		for maxDeque.Len() > 0 && nums[right] > nums[maxDeque.Back().Value.(int)] {
			maxDeque.Remove(maxDeque.Back())
		}
		maxDeque.PushBack(right)

		// Обновляем minDeque
		for minDeque.Len() > 0 && nums[right] < nums[minDeque.Back().Value.(int)] {
			minDeque.Remove(minDeque.Back())
		}
		minDeque.PushBack(right)

		// Проверяем условие
		for nums[maxDeque.Front().Value.(int)]-nums[minDeque.Front().Value.(int)] > 2 {
			// Сдвигаем левый указатель
			left++
			// Удаляем элементы из деков, если они вышли за пределы окна
			if maxDeque.Front().Value.(int) < left {
				maxDeque.Remove(maxDeque.Front())
			}
			if minDeque.Front().Value.(int) < left {
				minDeque.Remove(minDeque.Front())
			}
		}

		// Количество новых подмассивов, заканчивающихся на right
		count += int64(right - left + 1)
	}

	return count
}

func main() {
	// Пример входных данных
	nums := []int{1, 2, 3, 4, 5}

	// Вызов функции и вывод результата
	result := continuousSubarrays(nums)
	fmt.Println(result) // Ожидаемый вывод: 15
}
