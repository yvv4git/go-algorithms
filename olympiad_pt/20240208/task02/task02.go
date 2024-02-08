package main

import (
	"fmt"
	"strconv"
)

/*
Используя числа 1, 4, 5, 6, арифметические операции (сложение, вычитание, умножение, деление) и скобки, получить число 24.
Разрешается использовать только эти числа и только эти операции. Каждое число должно использоваться один и только один раз.
Операции и скобки можно использовать любое число раз. Нельзя объединять числа как цифры, составляя, например, AB, BCD и т.п.
*/

var numbers = []int{1, 4, 5, 6}
var target = 24

func evaluate(num1, num2 int, op rune) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		if num2 != 0 && num1%num2 == 0 {
			return num1 / num2
		}
		return -1 // Indicates division by zero or non-integer result
	}
	return -1 // Indicates an invalid operator
}

func solve(nums []int, target int, expression string) {
	if len(nums) == 1 && nums[0] == target {
		fmt.Println(expression + strconv.Itoa(nums[0]))
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for _, op := range []rune{'+', '-', '*', '/'} {
				result := evaluate(nums[i], nums[j], op)
				if result != -1 {
					newNums := make([]int, len(nums)-1)
					copy(newNums[:i], nums[:i])
					copy(newNums[i:], nums[i+1:j])
					copy(newNums[j-1:], nums[j:])
					newNums[i] = result

					newExpression := "(" + expression + strconv.Itoa(nums[i]) + string(op) + strconv.Itoa(nums[j]) + ")"
					solve(newNums, target, newExpression)
				}
			}
		}
	}
}

func main() {
	solve(numbers, target, "")
}
