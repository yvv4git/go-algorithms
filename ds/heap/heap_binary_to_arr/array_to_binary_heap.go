package main

// Как сконвертировать массив в Binary heap
func arrayToBinaryHeap(arr []int) (root *TreeNode) {
	arrLen := len(arr)

	var initNode func(curIdx int) *TreeNode

	initNode = func(curIdx int) *TreeNode {
		if curIdx > arrLen-1 {
			return nil
		}

		val := arr[curIdx]
		// Небольшое допущение, которое предполагает, что 0 - это null.
		// Просто не хотелось делать []*int, где можно использовать null.
		if val == 0 {
			return nil
		}

		leftIdx := leftNodeIdx(curIdx)
		rightIdx := rightNodeIdx(curIdx)

		return &TreeNode{
			Val:   val,
			Left:  initNode(leftIdx),
			Right: initNode(rightIdx),
		}
	}

	return initNode(0)
}

func leftNodeIdx(curIdx int) int {
	return (2 * curIdx) + 1
}

func rightNodeIdx(curIdx int) int {
	return (2 * curIdx) + 2
}
