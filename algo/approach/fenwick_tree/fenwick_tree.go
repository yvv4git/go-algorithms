package main

import "fmt"

// FenwickTree представляет дерево Фенвика.
type FenwickTree struct {
	tree []int // Массив для хранения дерева Фенвика.
}

// NewFenwickTree создает новое дерево Фенвика с заданным размером.
// Пространственная сложность: O(n), где n - размер дерева.
func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, size+1), // +1 потому что дерево Фенвика индексируется с 1.
	}
}

// Update обновляет значение в дереве Фенвика по заданному индексу.
// Временная сложность: O(log n), где n - размер дерева.
func (ft *FenwickTree) Update(index, value int) {
	for index < len(ft.tree) {
		ft.tree[index] += value
		index += index & -index // Переход к следующей вершине в дереве.
	}
}

// Query возвращает сумму элементов от начала массива до заданного индекса.
// Временная сложность: O(log n), где n - размер дерева.
func (ft *FenwickTree) Query(index int) int {
	sum := 0
	for index > 0 {
		sum += ft.tree[index]
		index -= index & -index // Переход к родительской вершине в дереве.
	}
	return sum
}

func main() {
	/*
		Этот пример демонстрирует базовое использование дерева Фенвика для вычисления префиксных сумм с эффективными обновлениями.
	*/
	// Создаем дерево Фенвика размером 10.
	fenwickTree := NewFenwickTree(10)

	// Обновляем значения в дереве.
	fenwickTree.Update(1, 3) // Добавляем 3 к элементу с индексом 1.
	fenwickTree.Update(2, 5) // Добавляем 5 к элементу с индексом 2.
	fenwickTree.Update(3, 2) // Добавляем 2 к элементу с индексом 3.

	// Запрашиваем сумму до определенных индексов.
	fmt.Println(fenwickTree.Query(1)) // Вывод: 3 (сумма до индекса 1)
	fmt.Println(fenwickTree.Query(2)) // Вывод: 8 (сумма до индекса 2)
	fmt.Println(fenwickTree.Query(3)) // Вывод: 10 (сумма до индекса 3)
}
