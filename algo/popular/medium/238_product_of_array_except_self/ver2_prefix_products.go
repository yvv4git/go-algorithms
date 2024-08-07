package main

func productExceptSelfV2(nums []int) []int {
	/*
		METHOD: Prefix products
		Данное решение задачи использует подход, который называется "префиксные произведения" (prefix products), но с оптимизацией по памяти.
		Вместо создания двух дополнительных массивов для хранения произведений элементов слева и справа от каждого элемента,
		оно использует только один дополнительный массив и одну переменную.

		TIME COMPLEXITY: O(n), где n — длина массива nums.
		Это объясняется тем, что мы проходим по массиву два раза: один раз для вычисления произведений элементов слева
		и один раз для вычисления произведений элементов справа. Каждый из этих проходов занимает линейное время,
		так как мы обрабатываем каждый элемент массива ровно один раз.

		SPACE COMPLEXITY: O(1), если не учитывать массив output как дополнительную память.
		Это объясняется тем, что мы используем только одну дополнительную переменную right для хранения произведения элементов справа.
		Массив output считается частью результата и не учитывается в дополнительной памяти. Если же учитывать массив output,
		то пространственная сложность будет O(n), так как мы создаем массив той же длины, что и входной массив nums.

	*/
	size := len(nums)
	right := 1
	output := make([]int, size, size)
	output[0] = 1
	for i := 1; i < size; i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	for i := size - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}