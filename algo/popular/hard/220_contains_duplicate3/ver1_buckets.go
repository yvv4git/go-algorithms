package main

import (
	"fmt"
)

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	/*
		METHOD: Buckets
		1. Использование корзин: Мы используем карту buckets для хранения чисел в "корзинах".
		Каждая корзина соответствует диапазону чисел размером valueDiff + 1. Это позволяет нам быстро проверять,
		есть ли число в текущей или соседних корзинах.

		2. Вычисление индекса корзины: Функция getBucketId вычисляет индекс корзины для заданного числа.
		Для отрицательных чисел мы корректируем индекс, чтобы они попадали в правильные корзины.

		3. Проверка условий: Для каждого числа мы проверяем, есть ли уже число в той же корзине или в соседних корзинах.
		Если есть, и разность между числами не превышает valueDiff, возвращаем true.

		4. Управление окном размера indexDiff: Мы удаляем корзины, которые выходят за пределы окна размера indexDiff,
		чтобы обеспечить, что разность индексов не превышает indexDiff.

		TIME COMPLEXITY: O(n), где n — количество элементов в массиве nums. Мы проходим по массиву один раз,
		и для каждого элемента выполняем операции с картой, которые занимают константное время.

		SPACE COMPLEXITY: O(min(n, indexDiff)), где n — количество элементов в массиве nums, а indexDiff — заданное число.
		В худшем случае мы храним indexDiff элементов в карте buckets.
	*/
	// Эта проверка if valueDiff < 0 { return false } используется для того, чтобы сразу вернуть false в случае, если valueDiff (разность между значениями элементов) меньше нуля.
	// Это может быть связано с особенностями задачи или с тем, как интерпретируется условие t.
	if valueDiff < 0 {
		return false
	}

	// Создаем карту для хранения значений в корзинах
	buckets := make(map[int]int)

	// Функция для вычисления индекса корзины
	getBucketId := func(num int) int {
		if num >= 0 {
			return num / (valueDiff + 1)
		}
		return (num / (valueDiff + 1)) - 1
	}

	for i, num := range nums {
		// Получаем индекс корзины для текущего числа
		id := getBucketId(num)

		// Проверяем, есть ли уже число в этой корзине
		if _, exists := buckets[id]; exists {
			return true
		}

		// Проверяем, есть ли число в соседних корзинах
		if prev, exists := buckets[id-1]; exists && num-prev <= valueDiff {
			return true
		}
		if next, exists := buckets[id+1]; exists && next-num <= valueDiff {
			return true
		}

		// Добавляем текущее число в соответствующую корзину
		buckets[id] = num

		// Удаляем корзину, которая выходит за пределы окна размера indexDiff
		if i >= indexDiff {
			delete(buckets, getBucketId(nums[i-indexDiff]))
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	t := 0
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t)) // Output: true

	nums = []int{1, 5, 9, 1, 5, 9}
	k = 2
	t = 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t)) // Output: false
}
