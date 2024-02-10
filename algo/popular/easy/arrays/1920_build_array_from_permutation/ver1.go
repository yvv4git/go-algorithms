package main

func buildArray(nums []int) []int {
	/*
		METHOD: Iteration
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве nums. Это обусловлено тем, что мы проходим по каждому элементу массива ровно один раз.
		SPACE COMPLEXITY: O(n), так как мы создаем новый массив ans, который имеет тот же размер, что и исходный массив nums.
	*/
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
