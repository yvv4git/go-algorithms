//unlint:unused
package basis

func twoSum(nums []int, target int) []int {
	for idx, number := range nums {
		for idx2, number2 := range nums {
			if number+number2 == target && idx != idx2 {
				return []int{idx, idx2}
			}
		}
	}

	return []int{}
}
