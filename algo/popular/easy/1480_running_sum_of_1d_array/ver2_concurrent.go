package main

import "sync"

func runningSumV2(nums []int) []int {
	/*
		METHOD: Concurrent
		- Инициализируем переменную `result` с длиной, равной длине массива `nums`.
		- Используем WaitGroup `wg` для ожидания завершения всех горутин.
		- Запускаем горутину для каждого элемента массива `nums`.
		- Внутри гор
		- Внутри горутины вычисляем сумму элементов от начала массива до текущего элемента.
		- Обновляем значение элемента `result[i]` путем прибавления значения `nums[i]`.
		- Используем Mutex `mu` для синхронизации доступа к переменной `result`.
		- После завершения всех горутин ожидаем завершения WaitGroup `wg`.
		- Возвращаем массив `result`.

		- TIME COMPLEXITY: O(n), где n - длина массива `nums`.
		- SPACE COMPLEXITY: O(n), так как используется дополнительный массив `result` длиной n.
	*/
	result := make([]int, len(nums))
	var wg sync.WaitGroup

	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			result[index] = result[index-1] + nums[index]
		}(i)
	}

	wg.Wait()

	return result
}
