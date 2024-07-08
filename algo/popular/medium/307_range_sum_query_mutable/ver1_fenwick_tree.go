package main

import "fmt"

// NumArray структура, которая будет использовать дерево Фенвика для быстрого вычисления сумм и обновления значений.
type NumArray struct {
	nums []int // Оригинальный массив
	tree []int // Дерево Фенвика
}

// Constructor конструктор для инициализации структуры NumArray.
func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1)
	for i := 0; i < n; i++ {
		updateTree(tree, i+1, nums[i])
	}
	return NumArray{nums, tree}
}

// Функция для обновления значения в дереве Фенвика.
// Временная сложность: O(log n), где n - количество элементов в массиве.
func updateTree(tree []int, index int, val int) {
	for index < len(tree) {
		tree[index] += val
		index += index & -index
	}
}

// Update метод для обновления значения в оригинальном массиве и дереве Фенвика.
// Временная сложность: O(log n), где n - количество элементов в массиве.
func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	updateTree(this.tree, index+1, diff)
}

// Функция для вычисления суммы первых index элементов в дереве Фенвика.
// Временная сложность: O(log n), где n - количество элементов в массиве.
func sumTree(tree []int, index int) int {
	sum := 0
	for index > 0 {
		sum += tree[index]
		index -= index & -index
	}
	return sum
}

// SumRange метод для вычисления суммы элементов от left до right включительно.
// Временная сложность: O(log n), где n - количество элементов в массиве.
func (this *NumArray) SumRange(left int, right int) int {
	return sumTree(this.tree, right+1) - sumTree(this.tree, left)
}

// Пример использования
func main() {
	/*
		METHOD: Fenwick Tree
		Мы используем дерево Фенвика (Fenwick Tree), которое позволяет эффективно выполнять операции обновления и вычисления суммы на отрезке.
		Дерево Фенвика представляет собой массив, где каждый элемент хранит частичную сумму элементов исходного массива.
		Благодаря особому способу индексации, операции обновления и вычисления суммы выполняются за логарифмическое время.

		TIME COMPLEXITY:
		- Constructor: O(n log n), где n - количество элементов в массиве.
		Это связано с тем, что для каждого элемента массива выполняется операция updateTree, которая имеет сложность O(log n).
		- Update: O(log n), так как в дереве Фенвика обновление значения требует обхода не более чем O(log n) узлов.
		- SumRange: O(log n), так как для вычисления суммы на отрезке требуется два вызова функции sumTree, каждый из которых имеет сложность O(log n).

		SPACE COMPLEXITY:
		- O(n), где n - количество элементов в массиве.
		Это связано с тем, что мы храним дополнительный массив tree размером n+1 для дерева Фенвика.

	*/
	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2)) // Вывод: 9
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2)) // Вывод: 8
}
