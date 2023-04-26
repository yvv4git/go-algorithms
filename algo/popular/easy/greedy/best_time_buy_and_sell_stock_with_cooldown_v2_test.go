package greedy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
# 309. Best Time to Buy and Sell Stock with Cooldown

## Medium
You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
- After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).


## Описание задачи (как я понял)
В общем, на вход подается массив целых чисел, которые представляют из себя стоимость актива по дням.
Надо купить актив дешевле, продать дороже.
НО есть ограничение - после продажи нельзя выполнить покупку в течении одного дня(cooldown day),


## Constraints:
- 1 <= prices.length <= 5000
- 0 <= prices[i] <= 1000


## Example 1:
Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]


## Example 2:
Input: prices = [1]
Output: 0
*/

func TestMaxProfitWithCoolDownV2(t *testing.T) {
	type args struct {
		prices []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				prices: []int{1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfitWithCoolDownV2(tt.args.prices)
			//t.Logf("res: %d", result)
			assert.Equal(t, tt.want, result)
		})
	}
}

func maxProfitWithCoolDownV2(prices []int) int {
	/*
		Какое-то сложное решение.
	*/
	return topDown(prices, 0, true, make(map[string]int))
}

func topDown(prices []int, idx int, isBuy bool, cache map[string]int) int {
	if idx >= len(prices) {
		return 0
	}
	key := fmt.Sprintf("%d:%t", idx, isBuy)
	if val, ok := cache[key]; ok {
		return val
	}

	if isBuy {
		buy := -prices[idx] + topDown(prices, idx+1, false, cache)
		skip := topDown(prices, idx+1, true, cache)
		cache[key] = max(buy, skip)

		return cache[key]
	}
	sell := prices[idx] + topDown(prices, idx+2, true, cache)
	skip := topDown(prices, idx+1, false, cache)
	cache[key] = max(sell, skip)

	return cache[key]
}
