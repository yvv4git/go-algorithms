package main

func canJumpV2(nums []int) bool {
	/*
		METHOD: Gready algorithm
		Подход, который используется в данном решении, называется "жадным алгоритмом" (greedy algorithm).
		Жадный алгоритм делает локально оптимальный выбор на каждом шаге с надеждой, что это приведет к глобальному оптимальному решению.
		В данном случае на каждом шаге мы выбираем максимально возможный индекс, который можно достичь, и проверяем, не достигли ли мы конца массива.

		Мы проходим по массиву и на каждом шаге обновляем максимальный индекс, который мы можем достичь на текущий момент.
		Если в какой-то момент текущий индекс превышает максимальный индекс, который мы можем достичь, то возвращаем false.
		Если же максимальный индекс, который мы можем достичь, становится больше или равен последнему индексу массива, то возвращаем

		TIME COMPLEXITY: O(n), где n — длина массива nums. Мы проходим по массиву один раз.

		SPACE COMPLEXITY: O(1), так как мы используем только константное количество дополнительной памяти.
	*/
	// Инициализируем переменную prev значением первого элемента массива.
	// Эта переменная будет хранить количество оставшихся шагов, которые мы можем сделать.
	prev := nums[0]

	// Начинаем цикл с индекса 1, так как первый элемент уже учтен в prev.
	for i := 1; i < len(nums); i++ {
		// Если prev становится 0, это означает, что мы не можем сделать ни одного шага дальше,
		// и поэтому возвращаем false.
		if prev == 0 {
			return false
		}

		// Обновляем prev, уменьшая его на 1 (так как мы делаем один шаг),
		// и берем максимум между этим уменьшенным значением и значением текущего элемента nums[i].
		// Это позволяет нам учесть возможность сделать больше шагов, если nums[i] больше, чем оставшиеся шаги.
		prev = max(prev-1, nums[i])
	}

	// Если мы успешно прошли весь массив, не встретив ситуации, когда prev становится 0,
	// то возвращаем true.
	return true
}