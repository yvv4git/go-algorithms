package main

func containsDuplicateV1(nums []int) bool {
	/*
		METHOD: Hash table, dictionary.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)

		Time complexity.
		Временная сложность функции containsDuplicateV1 составляет O(n), где n - количество элементов в массиве nums.
		Это связано с тем, что функция проходит по каждому элементу массива ровно один раз.

		SPACE COMPLEXITY:
		Пространственная сложность функции containsDuplicateV1 составляет O(n), где n - количество элементов в массиве nums.
		Это связано с тем, что в худшем случае (когда все элементы уникальны), функция будет хранить в словаре tmp все элементы массива nums.
	*/
	tmp := make(map[int]int)

	for _, v := range nums {

		tmp[v]++

		if tmp[v] > 1 {
			return true
		}
	}

	return false
}
