package greedy

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

/*
# 714. Best Time to Buy and Sell Stock with Transaction Fee

## Medium
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

## Описание задачи
Есть дни - это массив prices и значение - это цена актива в указанный день.
Fee - это платеж за каждую транзакцию.

## Constraints:
- 1 <= prices.length <= 5 * 104
- 1 <= prices[i] < 5 * 104
- 0 <= fee < 5 * 104

## Example 1:
Input: prices = [1,3,2,8,4,9], fee = 2
Output: 8
Explanation: The maximum profit can be achieved by:
- Buying at prices[0] = 1
- Selling at prices[3] = 8
- Buying at prices[4] = 4
- Selling at prices[5] = 9
The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

## Example 2:
Input: prices = [1,3,7,5,10,3], fee = 3
Output: 6

*/

func TestMaxProfitWithTransactionFreeWithBruteForce(t *testing.T) {
	type args struct {
		prices []int // Стоимость актива по дням
		fee    int   // Платеж за каждую транзакцию
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
		{
			name: "CASE-2",
			args: args{
				prices: []int{1, 3, 7, 5, 10, 3},
				fee:    3,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfitWithTransactionFeeV5(tt.args.prices, tt.args.fee)
			assert.Equal(t, tt.want, result)
		})
	}
}

// Общая функция для вычисления максимального из значений. Используется во всех вариантах.
func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func maxProfitWithTransactionFeeV1(prices []int, fee int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	cash := 0           // Сколько заработаем.
	sell := math.MinInt // Задаем минимальное значение, какое только возможно.

	for _, price := range prices {
		cash = max(cash, sell+price)
		sell = max(sell, cash-price-fee) // Считаем, сколько продали.
	}

	return cash
}

func maxProfitWithTransactionFeeV2(prices []int, fee int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	result, hold := 0, math.MinInt32

	for _, v := range prices {
		result, hold = max(result, hold+v), max(hold, result-v-fee)
	}

	return result
}

func maxProfitWithTransactionFeeV3(prices []int, fee int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	// Hold - сколько денег удержано.
	// Sold - сколько денег потрачено.
	hold, sold := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		hold, sold = max(hold, sold-prices[i]), max(sold, hold+prices[i]-fee)
	}
	return max(hold, sold)
}

func maxProfitWithTransactionFeeV4(prices []int, fee int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	length := len(prices)
	dp, curMax := make([]int, length+1), 0

	for i := length - 1; i >= 0; i-- {
		if i <= length-1 {
			curMax = max(curMax, (prices[i]+dp[i+1])-fee)
		}

		dp[i] = max(dp[i+1], curMax-prices[i])
	}

	return dp[0]
}

func maxProfitWithTransactionFeeV5(prices []int, fee int) int {
	/*
		Задача решается динамическим программированием.

		Time complexity: O(n)
		Space complexity: O(1)
	*/

	//bestIfOwn, bestIfDontOwn := -math.MaxInt32, 0
	bestIfOwn, bestIfDontOwn := -prices[0], 0

	for _, price := range prices {
		// Когда выгодно купить.
		if bestIfDontOwn-price > bestIfOwn {
			bestIfOwn = bestIfDontOwn - price
		}

		// Когда выгодно продать(bestIfOwn + price - fee).
		if bestIfOwn+price-fee > bestIfDontOwn {
			bestIfDontOwn = bestIfOwn + price - fee
		}
	}

	return bestIfDontOwn
}
