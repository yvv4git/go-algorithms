package main

import "fmt"

// Node - определение структуры узла QuadTree
type Node struct {
	Val         bool  // Значение узла (true, если 1, иначе false)
	IsLeaf      bool  // Флаг, указывающий, является ли узел листом
	TopLeft     *Node // Указатель на верхний левый дочерний узел
	TopRight    *Node // Указатель на верхний правый дочерний узел
	BottomLeft  *Node // Указатель на нижний левый дочерний узел
	BottomRight *Node // Указатель на нижний правый дочерний узел
}

// Основная функция для построения QuadTree из двумерного массива
func construct(grid [][]int) *Node {
	// Вызываем рекурсивную функцию для построения QuadTree, начиная с левого верхнего угла
	return buildQuadTree(grid, 0, 0, len(grid))
}

// Рекурсивная функция для построения QuadTree
func buildQuadTree(grid [][]int, x, y, size int) *Node {
	// Проверяем, является ли текущий квадрант листом
	if isLeaf(grid, x, y, size) {
		// Создаем листовой узел с соответствующим значением
		return &Node{
			Val:    grid[x][y] == 1,
			IsLeaf: true,
		}
	}

	// Если квадрант не является листом, разбиваем его на четыре части
	mid := size / 2
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     buildQuadTree(grid, x, y, mid),         // Рекурсивно строим QuadTree для верхнего левого квадранта
		TopRight:    buildQuadTree(grid, x, y+mid, mid),     // Рекурсивно строим QuadTree для верхнего правого квадранта
		BottomLeft:  buildQuadTree(grid, x+mid, y, mid),     // Рекурсивно строим QuadTree для нижнего левого квадранта
		BottomRight: buildQuadTree(grid, x+mid, y+mid, mid), // Рекурсивно строим QuadTree для нижнего правого квадранта
	}
}

// Функция для проверки, является ли квадрант листом
func isLeaf(grid [][]int, x, y, size int) bool {
	// Получаем значение первого элемента в квадранте
	val := grid[x][y]
	// Проверяем все элементы в квадранте
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			// Если найден элемент, отличный от val, квадрант не является листом
			if grid[i][j] != val {
				return false
			}
		}
	}
	// Если все элементы одинаковы, квадрант является листом
	return true
}

// Функция для печати QuadTree (для демонстрации)
func printQuadTree(node *Node, level int) {
	if node == nil {
		return
	}
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%sIsLeaf: %v, Val: %v\n", indent, node.IsLeaf, node.Val)
	printQuadTree(node.TopLeft, level+1)
	printQuadTree(node.TopRight, level+1)
	printQuadTree(node.BottomLeft, level+1)
	printQuadTree(node.BottomRight, level+1)
}

func main() {
	/*
		METHOD: Recursion
		Основная идея заключается в том, чтобы разбить двумерный массив на четыре квадранта и проверить,
		являются ли все элементы в каждом квадранте одинаковыми. Если да, то создается листовой узел (leaf node),
		иначе рекурсивно обрабатывается каждый квадрант.

		TIME COMPLEXITY: O(n^2 * log n)
		Проверка каждого квадранта занимает O(n^2) времени, и таких проверок O(log n) на каждом уровне рекурсии.

		SPACE COMPLEXITY: O(n^2)
		В худшем случае, каждый элемент массива может соответствовать отдельному узлу в QuadTree,
		что дает O(n^2) пространственной сложности. Также учитывается пространство стека вызовов,
		которое составляет O(log n), но оно менее значимо по сравнению с O(n^2).
	*/
	grid := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1},
	}

	root := construct(grid)
	printQuadTree(root, 0)
}
