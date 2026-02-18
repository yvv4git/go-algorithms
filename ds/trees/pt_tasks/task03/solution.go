package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// TreeNode представляет узел бинарного дерева поиска
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert добавляет новое значение в бинарное дерево поиска итеративно
// Если значение уже существует, оно не добавляется
//
// Time complexity: O(h), где h - высота дерева
// Worst case: O(n), в сбалансированном дереве: O(log n)
// Space complexity: O(1)
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
			// value == current.Value, элемент уже существует
			break
		}
	}

	return t
}

// InOrderTraversal выполняет симметричный обход дерева
// и добавляет значения в предоставленный срез
// Порядок обхода: левый -> корень -> правый
// Для BST это даёт значения в порядке возрастания
//
// Time complexity: O(n)
// Space complexity: O(h) для стека рекурсии, где h - высота дерева
func (t *TreeNode) InOrderTraversal(result *[]int) {
	if t == nil {
		return
	}
	t.Left.InOrderTraversal(result)
	*result = append(*result, t.Value)
	t.Right.InOrderTraversal(result)
}

func main() {
	// Читаем весь stdin целиком
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	var root *TreeNode
	start := 0

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

	// Выполняем симметричный обход для получения отсортированных значений
	var result []int
	root.InOrderTraversal(&result)

	// Выводим каждое значение на отдельной строке
	for _, v := range result {
		fmt.Println(v)
	}
}

// Time complexity всего алгоритма:
// - Построение дерева: O(n * h), где n - количество элементов, h - высота дерева
// - Симметричный обход: O(n)
// - Итого: O(n * h)
// - Worst case (отсортированный ввод): O(n²)
// - Средний случай (случайный ввод): O(n log n)
//
// Space complexity:
// - O(n) для хранения дерева
// - O(n) для хранения результатов
// - O(h) для стека рекурсии при обходе
