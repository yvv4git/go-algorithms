package main

func majorityElementHashMap(nums []int) int {
	/*
		METHOD: Hash map.
		0. Проходим по входному списку и заполняем хэш количество повторов каждого элемента.
		1. Затем проходим по хэши и определяем самый часто встречающийся элемент.

		Time complexity : O(n)
		Space complexity : O(n)
	*/

	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	max := 0
	candidate := 0
	for key, val := range hash {
		if max < val {
			max = val
			candidate = key
		}
	}

	return candidate
}
