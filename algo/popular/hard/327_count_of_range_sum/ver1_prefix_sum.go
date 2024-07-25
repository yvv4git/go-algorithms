package main

func countRangeSumV1(nums []int, lower int, upper int) int {
	// Создаем массив префиксных сумм
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Инициализируем счетчик
	var result int

	// Двойной цикл для перебора всех возможных подмассивов
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums)+1; j++ {
			// Вычисляем сумму текущего подмассива
			if sum := prefix[j] - prefix[i]; lower <= sum && sum <= upper {
				result++
			}
		}
	}

	return result
}
