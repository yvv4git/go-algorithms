package basis

func processing(idx int, arr []int) int {
	// Вот это место называется "База рекурсии".
	// Т.е. когда решение задачи можно выполнить без вызова самой себя.
	// Например, ниже не стоит снова вызывать рекурсию, т.к. выйдем за пределы массива.
	if idx == len(arr)-1 {
		return arr[idx]
	}

	maxEnd := processing(idx+1, arr)
	if arr[idx] > maxEnd {
		return arr[idx]
	} else {
		return maxEnd
	}
}

// MaxByRecursion - used recursion algorithm for find max element in array
func MaxByRecursion(arr []int) int {
	return processing(0, arr)
}