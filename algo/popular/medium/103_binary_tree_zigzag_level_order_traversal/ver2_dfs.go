package main

// Функция zigzagLevelOrderV2 принимает корень бинарного дерева и возвращает узлы дерева в "змейкой" порядке обхода.
// Если корень равен nil, то возвращается пустой массив.
func zigzagLevelOrderV2(root *TreeNode) [][]int {
	/*
		METHOD: DFS

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы последнего уровня дерева.
	*/

	// Проверка наличия корня дерева
	if root == nil {
		return [][]int{}
	}

	// Инициализация массива для хранения результатов обхода
	ans := [][]int{}
	// Вызов функции dfs для обхода в глубину
	dfs(root, 0, &ans)
	// Разворот элементов на нечетных уровнях
	for i := 0; i < len(ans); i++ {
		if i%2 == 0 {
			continue
		}
		reversePointer(&ans[i])
	}
	return ans
}

// Функция reversePointer принимает указатель на слайс и меняет порядок элементов в нем на обратный.
func reversePointer(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n/2; i++ {
		(*arr)[i], (*arr)[n-1-i] = (*arr)[n-1-i], (*arr)[i]
	}
}

// Функция dfs реализует обход в глубину для бинарного дерева.
// Она принимает узел, текущий уровень и указатель на массив для хранения результатов.
func dfs(node *TreeNode, level int, ans *[][]int) {
	// Если текущий уровень еще не добавлен в массив результатов, то он добавляется
	if level >= len(*ans) {
		*ans = append(*ans, []int{node.Val})
	} else {
		// Иначе значение узла добавляется в слайс на текущем уровне
		(*ans)[level] = append((*ans)[level], node.Val)
	}
	// Рекурсивный вызов dfs для левого и правого дочерних узлов, если они существуют
	if node.Left != nil {
		dfs(node.Left, level+1, ans)
	}
	if node.Right != nil {
		dfs(node.Right, level+1, ans)
	}
}
