package min_cost_climbing_stairs

/*
746. Min Cost Climbing Stairs


Task.
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.

Task RU.
Вам дается целочисленный массив cost, где cost[i] - это стоимость i-й ступени лестницы. Как только вы оплатите стоимость, вы сможете подняться либо на одну, либо на две ступеньки.
Вы можете начать либо с шага с индексом 0, либо с шага с индексом 1.
Верните минимальную стоимость, чтобы добраться до верха этажа.


Constraints:
2 <= cost.length <= 1000
0 <= cost[i] <= 999

Example 1:
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.

Example 2:
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.
*/

func minCostClimbingStairs(cost []int) int {
	/*
		Суть алгоритма - подниматься на 1 или 2 ступеньки с наименьшими затратами.
		Возникает ощущение, что это задача на динамическое программирование.
	*/
	n := len(cost)

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
