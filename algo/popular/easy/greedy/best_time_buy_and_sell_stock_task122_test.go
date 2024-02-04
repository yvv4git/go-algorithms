package greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
# 122. Best Time to Buy and Sell Stock II

## Easy
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
Find and return the maximum profit you can achieve.

## Описание задачи (как я понял)
В общем, на вход подается массив целых чисел, которые представляют из себя стоимость актива по дням.
Надо купить актив дешевле, продать дороже.

## Constraints:
1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104


## Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.

## Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.

## Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
*/

func TestMaxProfitTask2(t *testing.T) {
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
			want: 7,
		},
		{
			name: "CASE-2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "CASE-3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfitTask2(tt.args.prices)
			//t.Logf("res: %d", result)
			assert.Equal(t, tt.want, result)
		})
	}
}

func maxProfitTask2(prices []int) int {
	/*
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	var result int // Здесь храним результат.

	// Начнем со второго элемента и будем сравнивать его с предыдущим.
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] { // Если цена за i-ый день больше, чем цена за i-1 день, то можно покупать/продавать.
			result += prices[i] - prices[i-1] // Посчитаем полученную выходу.
		}
	}

	return result
}
