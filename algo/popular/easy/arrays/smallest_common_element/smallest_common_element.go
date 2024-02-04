package smallest_common_element

/*
 * Дана матрица, в которой каждая строка отсортирована в неубывающем порядке.
 * Требуется вернуть наименьший общий элемент во всех строках.
 * Если общего элемента нет, верните -1.
 */

func smallestCommonElement(m [][]int) int {
	/*
		TIME COMPLEXITY: O(M x N).
		Auxiliary Space: O(M)

		The steps are as follows:
		1. Iterate elements of the first row.
		2. For each element of the first row, iterate the remaining M - 1 rows and linearly search the element.
		3. If the element is not present in some row then move to the next element of the first row and repeat step 2.
		4. Return the element currently being searched if present in all rows, This is the smallest common element.
		5. If you reach the end of the first row, return -1.
	*/
	minElement := -1

	for _, v := range m[0] { // Итерируемся по элементам первой строки и ищем этот элемент в других строках бинарным поиском.
		for i := 1; i < len(m); i++ {
			if find := binarySearch(m[i], v); find == -1 {
				break
			} else {
				minElement = v
			}
		}
	}

	return minElement
}

// Используем бинарный поиск. Возвращает id элемента в массиве.
func binarySearch(list []int, target int) int {
	var low, max = 0, len(list) - 1
	if max < low {
		return -1
	}

	for low <= max {
		var mid = low + (max-low)/2
		if list[mid] == target {
			return mid
		}

		if target > list[mid] {
			low = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}
