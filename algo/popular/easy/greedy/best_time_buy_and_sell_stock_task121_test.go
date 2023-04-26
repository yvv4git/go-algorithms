package greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

# 121. Best Time to Buy and Sell Stock

## Easy
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

## Описание задачи
В общем, на вход подается массив целых чисел, которые представляют из себя стоимость актива по дням.
Надо купить актив дешевле, продать дороже.

## Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

## Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

## Constraints:
- 1 <= prices.length <= 105
- 0 <= prices[i] <= 104
*/

func TestMaxProfit(t *testing.T) {
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "CASE-2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit(tt.args.prices)
			t.Logf("res: %d", result)
			assert.Equal(t, tt.want, result)
		})
	}
}

func maxProfit(prices []int) (result int) {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	var minPrice = prices[0] // Предполагаем, что это минимальная цена для покупки.

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i] // Если же первая цена больше следующей, то покупать невыгодно. Будем покупать актив с меньшей стоимостью.
		} else if (prices[i] - minPrice) > result {
			result = prices[i] - minPrice
		}
	}

	return
}
