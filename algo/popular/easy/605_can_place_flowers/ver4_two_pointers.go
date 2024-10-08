package main

func canPlaceFlowersV4(flowerbed []int, n int) bool {
	/*
		METHOD: Two Pointers
		Данная функция использует два указателя для проверки, можно ли посадить n цветов на клумбу,
		представленную массивом flowerbed, соблюдая правило, что цветы не должны расти рядом друг с другом.
		Функция проходит по всем элементам массива и проверяет, можно ли посадить цветок на каждой пустой
		позиции (0). Если можно, то она помечает позицию как занятую (меняет 0 на 1) и увеличивает счетчик
		посаженных цветов. Если количество посаженных цветов достигает n, функция возвращает true.
		Если после прохода по массиву количество посаженных цветов меньше n, функция возвращает false.

		TIME COMPLEXITY: O(N)
		Временная сложность данной функции составляет O(N), где N — длина массива flowerbed. Это связано
		с тем, что функция выполняет один проход по всем элементам массива.

		SPACE COMPLEXITY: O(1)
		Пространственная сложность данной функции составляет O(1), так как она использует фиксированное
		количество дополнительной памяти (переменные count, length, i) вне зависимости от размера входного
		массива. Массив flowerbed модифицируется на месте, и дополнительная память для хранения данных
		не выделяется.
	*/
	count := 0               // Инициализируем счетчик для количества посаженных цветов
	length := len(flowerbed) // Получаем длину массива flowerbed
	i := 0                   // Инициализируем указатель для прохода по массиву

	// Проходим по всем позициям в массиве flowerbed, пока не посадим n цветов
	for i < length && count < n {
		// Проверяем, можно ли посадить цветок на текущей позиции
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			// Помечаем позицию как занятую
			flowerbed[i] = 1
			// Увеличиваем счетчик посаженных цветов
			count++
			// Пропускаем следующую позицию, так как она не может быть использована
			i += 2
		} else {
			// Переходим к следующей позиции
			i++
		}
	}

	// Возвращаем true, если количество посаженных цветов больше или равно n
	return count >= n
}
