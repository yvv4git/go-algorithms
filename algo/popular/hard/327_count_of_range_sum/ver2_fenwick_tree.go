package main

import "sort"

func countRangeSumV2(nums []int, lower int, upper int) int {
	/*
		METHOD: Fenwick Tree
		Я буду использовать дерево Фенвика для решения этой задачи.
		Дерево Фенвика (Fenwick Tree, также известное как Binary Indexed Tree) — это структура данных,
		которая позволяет эффективно выполнять операции обновления элементов и вычисления префиксных сумм в массиве.

		TIME COMPLEXITY: O(n log n), где n — длина массива nums.
		Это связано с сортировкой префиксных сумм и операцией обновления и запроса в дереве Фенвика,
		которые выполняются за O(log n).

		SPACE COMPLEXITY: O(n), где n — длина массива nums. Это связано с хранением префиксных сумм и дерева Фенвика.
	*/
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Сортируем префиксные суммы
	sortedPrefix := make([]int, len(prefix))
	copy(sortedPrefix, prefix)
	sort.Ints(sortedPrefix)

	// Создаем дерево Фенвика
	ft := NewFenwickTree(len(sortedPrefix))

	result := 0
	for _, p := range prefix {
		lowerBound := p - upper
		upperBound := p - lower
		result += ft.Query(indexOf(sortedPrefix, upperBound+1)) - ft.Query(indexOf(sortedPrefix, lowerBound))
		ft.Update(indexOf(sortedPrefix, p)+1, 1)
	}

	return result
}

// FenwickTree представляет дерево Фенвика
type FenwickTree struct {
	tree []int
}

// NewFenwickTree создает новое дерево Фенвика
func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, n+1),
	}
}

// Update обновляет дерево Фенвика
func (ft *FenwickTree) Update(i, delta int) {
	for i < len(ft.tree) {
		ft.tree[i] += delta
		i += i & -i
	}
}

// Query возвращает сумму элементов до индекса i
func (ft *FenwickTree) Query(i int) int {
	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= i & -i
	}
	return sum
}

// indexOf возвращает индекс элемента в отсортированном массиве
func indexOf(sorted []int, value int) int {
	return sort.SearchInts(sorted, value)
}
