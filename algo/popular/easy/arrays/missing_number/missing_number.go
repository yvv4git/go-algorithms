package missingnumber

func missingNumber(nums []int) int {
	/*
		Method: Arithmetic
		Time complexity: O(n)
		Space complexity: O(1)

		Задача решается элементарно!
		Т.е. у нас есть:
		- числа от 0 до n
		- числа уникальные, т.е. 0, 1, 2, 3, 4, 5..., n

		Например, было [0, 1, 2, 3], затем 2 удалили и осталось [0, 1, 3]. Как вычислить 2?
		Можно итерируясь по индексам посчитать, какая должна быть сумма чисел: ind += i+1, 0 + 1 + 2 + 3 = 6.
		Затем посчитать текущую сумму элементов: 0 + 1 + 3 = 4.
		Следовательно, какого числа не хватает? 6-4 = 2.
	*/
	var sum int
	var ind int

	for i, v := range nums {
		sum += v
		ind += i + 1
	}

	return ind - sum
}
