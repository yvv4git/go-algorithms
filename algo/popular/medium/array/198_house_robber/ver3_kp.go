package main

func robV3(nums []int) int {
	/*
		METHOD: Knapsack Problem = Dynamic programming
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого кода O(n), где n - количество домов.
		Это обусловлено циклом for i := 2; i < len(nums); i++, который проходит по каждому дому ровно один раз.

		Space complexity
		Пространственная сложность также O(n), так как мы используем дополнительный массив dp размером n для хранения результатов вычислений.
		Этот массив не учитывается в входных данных, поэтому считается постоянным пространств.
		Таким образом, временная и пространственная сложности этого кода составляют O(n).

		Подход к задаче "рюкзака" (Knapsack Problem) - это метод решения оптимизационной задачи,
		в которой ищется набор предметов, который максимизирует суммарную ценность или минимизирует вес,
		при этом не превышая заданный предел вместимости рюкзака.
	*/
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func main() {
//	// Пример использования
//	nums := []int{1, 2, 3, 1}
//	fmt.Println(rob(nums)) // Вывод: 4
//}
