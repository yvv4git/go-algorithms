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
var targetNumber = 48

func evaluate(x, y int, op rune) (int, bool) {
	switch op {
	case '+':
		return x + y, true
	case '-':
		return x - y, true
	case '*':
		return x * y, true
	case '/':
		if y != 0 && x%y == 0 {
			return x / y, true
		}
		return 0, false
	}
	return 0, false
}

func solveWithRecursion(nums []int, target int, expression string) {
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
					solveWithRecursion(newNums, target, newExpression)
				}
			}
		}
	}
}

func main() {
	solveWithRecursion(numbers, targetNumber, "")
}
