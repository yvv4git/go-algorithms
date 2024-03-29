package main

func dp(l, r int, memo *map[[2]int][]*TreeNode) []*TreeNode {
	/*
		METHOD: Dynamic programming with memoization
		Динамическое программирование — это метод решения сложных задач, который разбивает их на более простые подзадачи,
		решает эти подзадачи и сохраняет их решения для повторного использования.
		В этом случае, мы используем карту (map) для хранения уже вычисленных подзадач, что позволяет избежать повторных вычислений.
		Мемоизация здесь особенно эффективна, потому что многие подзадачи в этой задаче являются пересекающимися,
		то есть решение одной подзадачи может использоваться для решения другой.
		Например, когда мы решаем подзадачу для диапазона [1, 5], мы можем использовать уже вычисленные деревья для подзадач [1, 4] и [5, 5].


		TIME COMPLEXITY: O(4^n / n^(3/2)), где n - количество узлов в дереве.
		Это связано с тем, что для каждого узла мы генерируем все возможные левое и правое поддеревья,
		и количество таких комбинаций для каждого узла является каталаном, который растет экспоненциально с n.

		SPACE COMPLEXITY: O(4^n / n^(3/2)), так как для хранения всех уникальных деревьев потребуется этот объем памяти.
		Каждое дерево требует O(n) памяти для хранения узлов, и существует O(4^n / n^(3/2)) деревьев.

		Что за переменные:
		- l левая граница
		- r правая граница
		- memo запоминаем ответы к подзадачам
	*/
	if l > r {
		// нет потомков
		return append(make([]*TreeNode, 0), nil)
	}
	key := [2]int{l, r}
	if v, ok := (*memo)[key]; ok {
		// если эта подзадача уже решена
		return v
	}

	// решаем
	ret := make([]*TreeNode, 0)
	for i := l; i <= r; i++ { // каждое число в пределах границ
		leftPart := dp(l, i-1, memo)
		rightPart := dp(i+1, r, memo)
		for j := 0; j < len(leftPart); j++ {
			for k := 0; k < len(rightPart); k++ {
				// перебираем все комбинации
				root := &TreeNode{
					Val:   i, // i является корнем
					Left:  leftPart[j],
					Right: rightPart[k],
				}
				ret = append(ret, root)
			}
		}
	}
	(*memo)[key] = ret // запоминаем
	return ret
}

func generateTreesV3(n int) []*TreeNode {
	memo := make(map[[2]int][]*TreeNode, 0)
	return dp(1, n, &memo)
}
