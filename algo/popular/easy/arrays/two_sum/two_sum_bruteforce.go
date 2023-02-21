//unlint:unused
package main

func twoSumBruteforce(nums []int, target int) []int {
	/*
		Method: Bruteforce.
		Смысл данного метода в том, что мы просто перебираем все комбинации.

		Время работы: O(n^2)
		Память: O(1)
	*/
	for idx, number := range nums {
		for idx2, number2 := range nums {
			if number+number2 == target && idx != idx2 {
				return []int{idx, idx2}
			}
		}
	}

	return []int{}
}
