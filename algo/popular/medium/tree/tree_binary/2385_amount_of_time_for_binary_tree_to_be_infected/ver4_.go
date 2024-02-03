package main

// Функция amountOfTimeV4 находит время, необходимое для инфицирования всего дерева, начиная с узла start.
func amountOfTimeV4(root *TreeNode, start int) int {
	/*
		Method: BFS
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого кода составляет O(n), где n - количество узлов в дереве.
		Это связано с тем, что мы проходим по каждому узлу ровно один раз.

		Space complexity
		Пространственная сложность также составляет O(n),
		так как в худшем случае (когда дерево является сбалансированным) мы храним в словарях father и infected все узлы дерева.
	*/
	// Создаем словарь, где ключ - значение узла, а значение - родительский узел.
	father := make(map[int]*TreeNode)
	// Инициализируем указатель на узел, с которого начинается инфицирование.
	startNode := &(root)
	// Запускаем обход дерева, чтобы найти узел start и заполнить словарь father.
	walk(root, father, start, startNode)
	// Запускаем процесс инфицирования, начиная с узла start.
	return infect(*startNode, father)
}

// Функция walk выполняет обход дерева и заполняет словарь father.
func walk(node *TreeNode, father map[int]*TreeNode, start int, startNode **TreeNode) {
	// Если значение текущего узла равно start, то обновляем указатель на узел start.
	if node.Val == start {
		*startNode = node
	}
	// Если у текущего узла есть левый потомок, то добавляем его в словарь father и продолжаем обход.
	if node.Left != nil {
		father[node.Left.Val] = node
		walk(node.Left, father, start, startNode)
	}
	// Если у текущего узла есть правый потомок, то добавляем его в словарь father и продолжаем обход.
	if node.Right != nil {
		father[node.Right.Val] = node
		walk(node.Right, father, start, startNode)
	}
}

// Функция infect выполняет процесс инфицирования.
func infect(node *TreeNode, father map[int]*TreeNode) int {
	// Создаем словарь, где ключ - значение узла, а значение - флаг, указывающий, инфицирован ли узел.
	infected := make(map[int]bool)
	// Помечаем узел start как инфицированный.
	infected[node.Val] = true
	// Инициализируем список текущих узлов.
	currNodes := []*TreeNode{node}
	// Инициализируем счетчик шагов.
	step := -1
	// Пока есть узлы для обработки, выполняем следующие действия.
	for len(currNodes) > 0 {
		// Увеличиваем счетчик шагов.
		step++
		// Инициализируем список следующих узлов.
		nextNodes := make([]*TreeNode, 0, len(currNodes))
		// Для каждого текущего узла выполняем следующие действия.
		for _, n := range currNodes {
			// Если у текущего узла есть левый потомок и он не инфицирован, то помечаем его как инфицированный и добавляем в список следующих узлов.
			if n.Left != nil && infected[n.Left.Val] == false {
				infected[n.Left.Val] = true
				nextNodes = append(nextNodes, n.Left)
			}
			// Если у текущего узла есть правый потомок и он не инфицирован, то помечаем его как инфицированный и добавляем в список следующих узлов.
			if n.Right != nil && infected[n.Right.Val] == false {
				infected[n.Right.Val] = true
				nextNodes = append(nextNodes, n.Right)
			}
			// Если у текущего узла есть родительский узел и он не инфицирован, то помечаем его как инфицированный и добавляем в список следующих узлов.
			if father[n.Val] != nil && infected[father[n.Val].Val] == false {
				infected[father[n.Val].Val] = true
				nextNodes = append(nextNodes, father[n.Val])
			}
		}
		// Обновляем список текущих узлов.
		currNodes = nextNodes
	}
	// Возвращаем количество шагов.
	return step
}
