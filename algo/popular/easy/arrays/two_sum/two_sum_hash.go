package main

func twoSumHash(nums []int, target int) []int {
	/*
		METHOD: Hash table.

		Время работы: O(n)
		Память: O(n)
	*/
	hash := make(map[int]int)

	for idx, num := range nums {
		numToFind := target - num

		if idx2, ok := hash[numToFind]; ok {
			return []int{idx, idx2}
		}

		hash[num] = idx
	}

	return []int{}
}
