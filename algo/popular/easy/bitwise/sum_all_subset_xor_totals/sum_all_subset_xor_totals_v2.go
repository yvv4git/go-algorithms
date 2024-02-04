package sum_all_subset_xor_totals

func subsetXORSumV2(nums []int) int {
	/*
		METHOD: Backtrack with DFS(recursion)
		TIME COMPLEXITY: O(e+v)
		Space complexity: O(1)
	*/
	return dfsBacktrack(0, nums, 0)
}

func dfsBacktrack(idx int, nums []int, curXor int) int {
	if len(nums) == 0 || idx == len(nums) {
		return curXor
	}

	withNum := dfsBacktrack(idx+1, nums, curXor^nums[idx])
	withoutNum := dfsBacktrack(idx+1, nums, curXor)

	return withNum + withoutNum
}
