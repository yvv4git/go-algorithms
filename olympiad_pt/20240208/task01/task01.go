package main

import (
	"fmt"
	"strconv"
)

/*
Используя числа 2, 5, 3, 8, арифметические операции (сложение, вычитание, умножение, деление) и скобки, получить число 16.
Разрешается использовать только эти числа и только эти операции. Каждое число должно использоваться один и только один раз.
Операции и скобки можно использовать любое число раз. Нельзя объединять числа как цифры, составляя, например, AB, BCD и т.п.
*/

var numbers = []int{2, 5, 3, 8}
var target = 16

func solve(nums []int, target int, expression string) {
	if len(nums) == 1 && nums[0] == target {
		fmt.Println(expression + strconv.Itoa(nums[0]))
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			newNums := make([]int, len(nums)-1)
			copy(newNums[:i], nums[:i])
			copy(newNums[i:], nums[i+1:j])
			copy(newNums[j-1:], nums[j:])

			newExpression := expression + "(" + strconv.Itoa(nums[i]) + "+" + strconv.Itoa(nums[j]) + ")"
			newNums[i] = nums[i] + nums[j]
			solve(newNums, target, newExpression)

			newExpression = expression + "(" + strconv.Itoa(nums[i]) + "-" + strconv.Itoa(nums[j]) + ")"
			newNums[i] = nums[i] - nums[j]
			solve(newNums, target, newExpression)

			newExpression = expression + "(" + strconv.Itoa(nums[i]) + "*" + strconv.Itoa(nums[j]) + ")"
			newNums[i] = nums[i] * nums[j]
			solve(newNums, target, newExpression)

			if nums[j] != 0 {
				newExpression = expression + "(" + strconv.Itoa(nums[i]) + "/" + strconv.Itoa(nums[j]) + ")"
				newNums[i] = nums[i] / nums[j]
				solve(newNums, target, newExpression)
			}
		}
	}
}

func main() {
	solve(numbers, target, "")
}
