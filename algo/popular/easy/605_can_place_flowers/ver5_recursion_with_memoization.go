package main

func canPlaceFlowersV5(flowerbed []int, n int) bool {
	/*
		METHOD: Recursion with Memoization
		Данная функция использует рекурсию с мемоизацией для проверки, можно ли посадить n цветов на клумбу,
		представленную массивом flowerbed, соблюдая правило, что цветы не должны расти рядом друг с другом.
		Функция canPlaceFlowersHelper рекурсивно проверяет каждую позицию в массиве и использует мемоизацию
		для хранения результатов уже вычисленных подзадач. Если количество посаженных цветов достигает n,
		функция возвращает true, иначе — false.

		TIME COMPLEXITY: O(N)
		Временная сложность данной функции составляет O(N), где N — длина массива flowerbed. Это связано
		с тем, что каждая позиция в массиве будет обработана не более одного раза благодаря мемоизации.

		SPACE COMPLEXITY: O(N)
		Пространственная сложность данной функции составляет O(N), так как она использует дополнительный
		словарь memo размером до N для хранения результатов уже вычисленных подзадач. Также учитывается
		пространство, используемое для стека вызовов рекурсии.
	*/
	memo := make(map[int]int)                                // Создаем словарь для мемоизации результатов
	return canPlaceFlowersHelper(flowerbed, n, 0, memo) >= n // Вызываем вспомогательную функцию
}

func canPlaceFlowersHelper(flowerbed []int, n, index int, memo map[int]int) int {
	if n == 0 { // Если нужно посадить 0 цветов, возвращаем 0
		return 0
	}
	if index >= len(flowerbed) { // Если индекс выходит за пределы массива, возвращаем 0
		return 0
	}
	if val, ok := memo[index]; ok { // Если результат для текущего индекса уже вычислен, возвращаем его
		return val
	}

	count := 0
	if flowerbed[index] == 0 && (index == 0 || flowerbed[index-1] == 0) && (index == len(flowerbed)-1 || flowerbed[index+1] == 0) {
		// Если можно посадить цветок на текущей позиции
		flowerbed[index] = 1                                             // Помечаем позицию как занятую
		count = 1 + canPlaceFlowersHelper(flowerbed, n-1, index+2, memo) // Рекурсивно вызываем функцию для следующей позиции
		flowerbed[index] = 0                                             // Возвращаем позицию в исходное состояние
	}

	// Выбираем максимальное количество цветов, которое можно посадить
	count = max(count, canPlaceFlowersHelper(flowerbed, n, index+1, memo))
	memo[index] = count // Сохраняем результат в словарь для мемоизации
	return count
}

//// Функция для нахождения максимума из двух чисел
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
