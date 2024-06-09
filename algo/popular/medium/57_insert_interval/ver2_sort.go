package main

import "sort"

func insertV2(intervals [][]int, new []int) [][]int {
	/*
		METHOD: Iterative + sort

		TIME COMPLEXITY: O(log n), где n - количество интервалов в массиве intervals.
		Это связано с использованием бинарного поиска для поиска позиций вставки нового интервала, который выполняется за логарифмическое время.

		SPACE COMPLEXITY: O(n), так как в худшем случае может потребоваться хранить все исходные интервалы в результирующем массиве intervals.
		Однако, в целом, так как мы используем дополнительную память только для результирующего массива,
		пространственная сложность составляет O(n), а не O(n + m), где m - количество дополнительных переменных.
	*/
	n := len(intervals)
	i := sort.Search(n, func(i int) bool { return intervals[i][0] > new[0] })
	j := sort.Search(n, func(j int) bool { return intervals[j][1] > new[1] })

	if i >= 1 && new[0] <= intervals[i-1][1] {
		new[0] = intervals[i-1][0]
		i--
	}

	if j < n && new[1] >= intervals[j][0] {
		new[1] = intervals[j][1]
		j++
	}

	return append(intervals[:i], append([][]int{new}, intervals[j:]...)...)
}
