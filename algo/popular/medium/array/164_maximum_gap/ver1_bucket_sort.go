package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	/*
		METHOD: Bucket sort
		Бакетная сортировка (Bucket Sort) помогает решить задачу "Maximum Gap" за счет группировки чисел в "бакеты" или интервалы,
		основываясь на их значениях. Это позволяет нам рассматривать только максимальное и минимальное значения в каждом бакете,
		что упрощает поиск максимального расстояния между элементами.

		Она работает путем разбиения входного массива на несколько отдельных "бакетов" или "контейнеров",
		каждый из которых содержит элементы из определенного диапазона, и затем сортирует каждый бакет отдельно.

		В этом коде мы используем бакетную сортировку для группировки чисел в бакеты в соответствии с их значениями.
		Затем мы находим максимальное расстояние между соседними бакетами,
		где минимальное значение одного бакета является максимальным значением предыдущего бакета.

		Таким образом, бакетная сортировка позволяет нам решить задачу "Maximum Gap" за линейное время O(n),
		где n - количество элементов в массиве, и использует линейную дополнительную память O(m), где m - количество бакетов.
		Это эффективнее, чем полную сортировку, которая имеет сложность O(n log n).

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве, так как мы проходим по массиву дважды:
		один раз для поиска минимального и максимального значений, а второй раз для поиска максимального расстояния между элементами.

		SPACE COMPLEXITY: O(m), где m - количество бакетов, которые мы создаем для хранения минимальных
		и максимальных значений элементов в каждом бакете. В худшем случае, когда все элементы различны,
		m будет равно n, но в среднем он будет меньше n.
	*/
	if len(nums) < 2 {
		return 0
	}

	// Находим минимальное и максимальное значения в массиве
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
	}

	// Вычисляем размер бакета и количество бакетов
	bucketSize := max(1, (maxNum-minNum)/(len(nums)-1))
	bucketNum := (maxNum-minNum)/bucketSize + 1

	// Создаем бакеты
	buckets := make([][]int, bucketNum)
	for i := range buckets {
		buckets[i] = []int{math.MaxInt64, math.MinInt64}
	}

	// Распределяем числа по бакетам
	for _, num := range nums {
		bucketIndex := (num - minNum) / bucketSize
		buckets[bucketIndex][0] = min(buckets[bucketIndex][0], num)
		buckets[bucketIndex][1] = max(buckets[bucketIndex][1], num)
	}

	// Находим максимальное расстояние между бакетами
	prevBucketMax := minNum
	maxGap := 0
	for _, bucket := range buckets {
		if bucket[0] == math.MaxInt64 {
			// Бакет пуст
			continue
		}
		maxGap = max(maxGap, bucket[0]-prevBucketMax)
		prevBucketMax = bucket[1]
	}

	return maxGap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(nums)) // Вывод: 3
}
