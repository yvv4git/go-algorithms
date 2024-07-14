package main

import "fmt"

// Функция productExceptSelf принимает массив целых чисел nums и возвращает новый массив,
// где каждый элемент является произведением всех элементов исходного массива, кроме самого себя.
func productExceptSelf(nums []int) []int {
	/*
		METHOD: Prefix products
		Используемый подход называется "префиксные произведения" (prefix products).
		Этот метод заключается в создании двух дополнительных массивов:
		1. left_products, который содержит произведения всех элементов слева от текущего элемента.
		2. right_products, который содержит произведения всех элементов справа от текущего элемента.

		Затем, для каждого элемента в исходном массиве, результат вычисляется как произведение соответствующих элементов из left_products и right_products.
		Этот подход позволяет решить задачу за линейное время и с линейным использованием дополнительной памяти.

		TIME COMPLEXITY: O(n). Мы проходим по массиву три раза: один раз для заполнения left_products,
		один раз для заполнения right_products и один раз для вычисления результата. Каждый проход занимает O(n) времени.

		SPACE COMPLEXITY: O(n). Мы используем два дополнительных массива left_products и right_products,
		каждый из которых имеет длину n. Также мы используем массив result длиной n для хранения результата.
	*/
	n := len(nums)
	// Инициализируем результирующий массив с длиной, равной длине исходного массива.
	result := make([]int, n)

	// Инициализируем массив left_products, где left_products[i] будет содержать
	// произведение всех элементов слева от nums[i].
	left_products := make([]int, n)
	left_products[0] = 1
	for i := 1; i < n; i++ {
		left_products[i] = left_products[i-1] * nums[i-1]
	}

	// Инициализируем массив right_products, где right_products[i] будет содержать
	// произведение всех элементов справа от nums[i].
	right_products := make([]int, n)
	right_products[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right_products[i] = right_products[i+1] * nums[i+1]
	}

	// Теперь для каждого индекса i в исходном массиве вычисляем результат
	// как произведение left_products[i] и right_products[i].
	for i := 0; i < n; i++ {
		result[i] = left_products[i] * right_products[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Вывод: [24, 12, 8, 6]
}
