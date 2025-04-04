package main

func kLengthApart(nums []int, k int) bool {
	/*
		APPROACH: Sliding Window
		1. Инициализируем предыдущую позицию единицы очень маленьким числом, чтобы первая единица всегда проходила проверку
		2. Проходим по nums
		3. Если текущий элемент равен 1, проверяем расстояние между текущей и предыдущей единицей
		4. Если расстояние между текущей и предыдущей единицей меньше k, возвращаем false
		5. Если расстояние между текущей и предыдущей единицей больше или равно k, обновляем позицию предыдущей единицы
		6. Если все проверки пройдены, возвращаем true

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(1)
		- Мы не используем дополнительной памяти, так как мы обновляем только одну переменную в каждом шаге.
	*/

	// Инициализируем предыдущую позицию единицы очень маленьким числом,
	// чтобы первая единица всегда проходила проверку
	prev := -1 << 31

	for i, num := range nums {
		if num == 1 {
			// Проверяем расстояние между текущей и предыдущей единицей
			// i - prev - 1 - это количество элементов между ними
			if i-prev-1 < k {
				return false
			}
			// Обновляем позицию предыдущей единицы
			prev = i
		}
	}

	// Если все проверки пройдены, возвращаем true
	return true
}

func main() {
	nums1 := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k1 := 2
	println(kLengthApart(nums1, k1)) // true

	nums2 := []int{1, 0, 0, 1, 0, 1}
	k2 := 2
	println(kLengthApart(nums2, k2)) // false
}
