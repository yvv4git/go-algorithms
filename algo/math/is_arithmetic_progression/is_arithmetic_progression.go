package is_arithmetic_progression

import "sort"

// CanMakeArithmeticProgression - used for check whether the array is an arithmetic progression.
func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	var dif = arr[1] - arr[0]

	for i := 1; i < len(arr); i++ {
		if dif != arr[i]-arr[i-1] {
			return false
		}
	}

	return true
}
