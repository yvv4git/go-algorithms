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

// Height возвращает высоту дерева итеративно (BFS)
func (t *TreeNode) Height() int {
	if t == nil {
		return 0
	}

	height := 0
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		levelSize := len(queue)
		height++

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return height
}

func main() {
	// Читаем весь stdin целиком
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(0)
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

	// Выводим высоту дерева
	fmt.Println(root.Height())
}
