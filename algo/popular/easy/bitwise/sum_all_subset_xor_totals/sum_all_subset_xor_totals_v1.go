package sum_all_subset_xor_totals

func subsetXORSumV1(nums []int) int {
	/*
		METHOD: Bruteforce with DFS(recursion)
		Time complexity: O(e+v)
		Space complexity: O(1)
	*/
	res := 0
	dfs(nums, 0, &res)
	return res
}

func dfs(nums []int, cur int, res *int) {
	if len(nums) == 0 {
		*res += cur
		return
	}

	dfs(nums[1:], cur^nums[0], res)
	dfs(nums[1:], cur, res)
}
