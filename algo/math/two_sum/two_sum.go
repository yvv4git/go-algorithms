package two_sum

// TwoSum - находит два элемента в массиве, сумма которых равна target. Сложность: O(n^2) - квадратичная.
func TwoSum(nums []int, target int) []int {
	for idx, element := range nums {
		for idx2, element2 := range nums {
			if (element+element2 == target) && (idx != idx2) {
				return []int{idx, idx2}
			}
		}
	}

	return []int{0, 0}
}
