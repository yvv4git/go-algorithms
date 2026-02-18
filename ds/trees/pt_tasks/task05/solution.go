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

// IsFork проверяет, является ли узел развилкой.
// Развилка — это узел, имеющий ровно два дочерних элемента (и левый, и правый).
//
// Временная сложность: O(1).
// Пространственная сложность: O(1).
func (t *TreeNode) IsFork() bool {
	return t != nil && t.Left != nil && t.Right != nil
}

// CollectForks собирает все развилки дерева в порядке возрастания.
// Использует симметричный обход (in-order) для получения отсортированного порядка.
//
// Временная сложность: O(n), где n — количество узлов в дереве.
// Пространственная сложность: O(h) для стека рекурсии, где h — высота дерева.
func (t *TreeNode) CollectForks(result *[]int) {
	if t == nil {
		return
	}

	// Сначала обходим левое поддерево.
	t.Left.CollectForks(result)

	// Если текущий узел — развилка, добавляем его в результат.
	if t.IsFork() {
		*result = append(*result, t.Value)
	}

	// Затем обходим правое поддерево.
	t.Right.CollectForks(result)
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

	// Собираем все развилки в порядке возрастания.
	var forks []int
	root.CollectForks(&forks)

	// Выводим каждую развилку на отдельной строке.
	for _, v := range forks {
		fmt.Println(v)
	}
}

// Временная сложность всего алгоритма:
// - Построение дерева: O(n * h), где n — количество элементов, h — высота дерева.
// - Сбор развилок: O(n).
// - Итого: O(n * h).
// - Худший случай (отсортированный ввод): O(n^2).
// - Средний случай (случайный ввод): O(n log n).
//
// Пространственная сложность:
// - O(n) для хранения дерева.
// - O(k) для хранения развилок, где k — количество развилок.
// - O(h) для стека рекурсии при обходе.
