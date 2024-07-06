package main

type finder struct {
	s []*TreeNode
}

func (f *finder) findNode(root, n *TreeNode) bool {
	if root == nil {
		return false
	} else if root.Val == n.Val {
		return true
	}

	if f.findNode(root.Left, n) {
		f.s = append(f.s, root.Left)
		return true
	} else if f.findNode(root.Right, n) {
		f.s = append(f.s, root.Right)
		return true
	}
	return false
}

func invArr(arr []*TreeNode) []*TreeNode {
	l := 0
	r := len(arr) - 1
	for l < r {
		tmp := arr[l]
		arr[l] = arr[r]
		arr[r] = tmp
		l++
		r--
	}
	return arr
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	/*
		METHOD: Two stack
		Подход "two stack" для нахождения наименьшего общего предка (LCA) в бинарном дереве заключается в использовании двух стеков
		для хранения путей от корня дерева до двух узлов, для которых ищется LCA. Основные шаги этого подхода можно описать следующим образом:
		1. Поиск путей до узлов:
		Для каждого из узлов p и q выполняется поиск пути от корня дерева до этого узла.
		Этот поиск может быть реализован с использованием рекурсии или итерации.
		В процессе поиска путь сохраняется в стеке.

		2. Инвертирование путей:
		Поскольку пути хранятся в стеках, они будут начинаться с узлов p и q и заканчиваться корнем дерева.
		Для удобства сравнения путей их необходимо инвертировать, чтобы они начинались с корня и заканчивались узлами p и q.

		3. Сравнение путей:
		После инвертирования путей они сравниваются элемент за элементом.
		Первый элемент, который не совпадает, указывает на то, что предыдущий элемент был последним общим предком для узлов p и q.

		4. Возврат LCA:
		Если все элементы путей совпадают до конца более короткого пути, то последний совпадающий элемент является LCA.

		TIME COMPLEXITY: O(n) - где n — количество узлов в дереве.
		Мы обходим дерево дважды для поиска путей до узлов p и q,
		а затем сравниваем пути для нахождения LCA.

		SPACE COMPLEXITY: O(n) - где n — количество узлов в дереве.
		Мы храним пути от корня до узлов p и q в массивах,
		что требует дополнительной памяти.
	*/
	if root == nil {
		return nil
	}
	f1 := &finder{[]*TreeNode{}}
	f2 := &finder{[]*TreeNode{}}
	f1.findNode(root, p)
	f2.findNode(root, q)
	f1.s = append(f1.s, root)
	f2.s = append(f2.s, root)
	a1 := invArr(f1.s)
	a2 := invArr(f2.s)
	n := minInt(len(a1), len(a2))
	for i := 0; i < n; i++ {
		if a1[i] != a2[i] {
			return a1[i-1]
		}
	}
	return a1[n-1]
}
