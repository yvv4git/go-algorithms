package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// TreeNode представляет узел бинарного дерева поиска.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert добавляет новое значение в бинарное дерево поиска итеративно.
// Если значение уже существует, оно не добавляется.
//
// Временная сложность: O(h), где h — высота дерева.
// Худший случай: O(n), в сбалансированном дереве: O(log n).
// Пространственная сложность: O(1).
func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return &TreeNode{Value: value}
	}

	current := t
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = &TreeNode{Value: value}
				break
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &TreeNode{Value: value}
				break
			}
			current = current.Right
		} else {
			// value == current.Value, элемент уже существует.
			break
		}
	}

	return t
}

// IsLeaf проверяет, является ли узел листом.
// Лист — это узел без дочерних элементов.
//
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func (t *TreeNode) IsLeaf() bool {
	return t != nil && t.Left == nil && t.Right == nil
}

// CollectLeaves собирает все листья дерева в порядке возрастания.
// Использует симметричный обход (in-order) для получения отсортированного порядка.
//
// Временная сложность: O(n), где n — количество узлов в дереве.
// Пространственная сложность: O(h) для стека рекурсии, где h — высота дерева.
func (t *TreeNode) CollectLeaves(result *[]int) {
	if t == nil {
		return
	}

	// Сначала обходим левое поддерево.
	t.Left.CollectLeaves(result)

	// Если текущий узел — лист, добавляем его в результат.
	if t.IsLeaf() {
		*result = append(*result, t.Value)
	}

	// Затем обходим правое поддерево.
	t.Right.CollectLeaves(result)
}

func main() {
	// Читаем весь stdin целиком.
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	var root *TreeNode
	start := 0

	// Парсим числа из входных данных.
	for i := 0; i <= len(data); i++ {
		if i == len(data) || data[i] == ' ' || data[i] == '\t' || data[i] == '\n' || data[i] == '\r' {
			if i > start {
				num, err := strconv.Atoi(string(data[start:i]))
				if err == nil {
					if num == 0 {
						break
					}
					root = root.Insert(num)
				}
			}
			start = i + 1
		}
	}

	// Собираем все листья в порядке возрастания.
	var leaves []int
	root.CollectLeaves(&leaves)

	// Выводим каждый лист на отдельной строке.
	for _, v := range leaves {
		fmt.Println(v)
	}
}

// Временная сложность всего алгоритма:
// - Построение дерева: O(n * h), где n — количество элементов, h — высота дерева.
// - Сбор листьев: O(n).
// - Итого: O(n * h).
// - Худший случай (отсортированный ввод): O(n^2).
// - Средний случай (случайный ввод): O(n log n).
//
// Пространственная сложность:
// - O(n) для хранения дерева.
// - O(k) для хранения листьев, где k — количество листьев.
// - O(h) для стека рекурсии при обходе.
