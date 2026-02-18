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

// InsertDepth добавляет новое значение в бинарное дерево поиска итеративно
// Возвращает глубину добавленного элемента и признак добавления
// Если значение уже существует, возвращает 0, false
//
// Временная сложность: O(h), где h - высота дерева
// В худшем случае O(n), в сбалансированном дереве O(log n)
// Пространственная сложность: O(1)
func (t *TreeNode) InsertDepth(value int) (depth int, inserted bool) {
	if t == nil {
		return 0, false // корень обрабатывается отдельно
	}

	depth = 1
	current := t

	for {
		depth++
		if value < current.Value {
			if current.Left == nil {
				current.Left = &TreeNode{Value: value}
				return depth, true
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &TreeNode{Value: value}
				return depth, true
			}
			current = current.Right
		} else {
			// value == current.Value, элемент уже существует
			return 0, false
		}
	}
}

func main() {
	// Читаем весь stdin целиком
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	var root *TreeNode
	var depths []int
	start := 0

	for i := 0; i <= len(data); i++ {
		if i == len(data) || data[i] == ' ' || data[i] == '\t' || data[i] == '\n' || data[i] == '\r' {
			if i > start {
				num, err := strconv.Atoi(string(data[start:i]))
				if err == nil {
					if num == 0 {
						break
					}
					if root == nil {
						root = &TreeNode{Value: num}
						depths = append(depths, 1)
					} else {
						if depth, inserted := root.InsertDepth(num); inserted {
							depths = append(depths, depth)
						}
					}
				}
			}
			start = i + 1
		}
	}

	// Выводим глубины добавленных элементов
	for i, d := range depths {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(d)
	}
	if len(depths) > 0 {
		fmt.Println()
	}
}

// Временная сложность всего алгоритма:
// - O(n * h), где n - количество элементов, h - высота дерева
// - В худшем случае (отсортированный ввод): O(n²)
// - В среднем случае (случайный ввод): O(n log n)
//
// Пространственная сложность:
// - O(n) для хранения дерева
// - O(n) для хранения результатов (можно выводить сразу, тогда O(1) дополнительно)
