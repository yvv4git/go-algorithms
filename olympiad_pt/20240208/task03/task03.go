package main

import (
	"fmt"
	"strconv"
)

/*
Используя числа 1, 6, 8, 9, арифметические операции (сложение, вычитание, умножение, деление) и скобки, получить число 48.
Разрешается использовать только эти числа и только эти операции. Каждое число должно использоваться один и только один раз.
Операции и скобки можно использовать любое число раз. Нельзя объединять числа как цифры, составляя, например, AB, BCD и т.п.
*/

var numbers = []int{1, 6, 8, 9}
var target = 48

func evaluate(num1, num2 int, op rune) (int, bool) {
	switch op {
	case '+':
		return num1 + num2, true
	case '-':
		return num1 - num2, true
	case '*':
		return num1 * num2, true
	case '/':
		if num2 != 0 && num1%num2 == 0 {
			return num1 / num2, true
		}
		return 0, false // Indicates division by zero or non-integer result
	}
	return 0, false // Indicates an invalid operator
}

func solve(nums []int, target int, expression string) {
	if len(nums) == 1 && nums[0] == target {
		fmt.Println(expression + strconv.Itoa(nums[0]))
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for _, op := range []rune{'+', '-', '*', '/'} {
				result, valid := evaluate(nums[i], nums[j], op)
				if valid {
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
