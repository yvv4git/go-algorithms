package main

// Функция rangeBitwiseAndV2 вычисляет побитовое И для всех чисел в диапазоне от left до right.
// Алгоритм основан на побитовом сдвиге и проверке каждого бита результата.
func rangeBitwiseAndV2(left int, right int) int {
	/*
		METHOD: Bitwise
		TIME COMPLEXITY: O(1) операций, так как внутри цикла выполняется фиксированное количество операций, не зависящих от размера входных данных.
		SPACE COMPLEXITY: O(1), так как используется фиксированное количество переменных, не зависящих от размера входных данных.
	*/
	// Находим начальный результат как побитовое И для левого и правого края диапазона
	result := left & right

	// Вычисляем расстояние между правой и левой границами диапазона
	distance := right - left + 1

	// Проходим по каждому биту результата
	for i := 0; i < 32; i++ { // За счет того, что количество операций ограничено 32 битами временная сложность равна O(1), т.е. константная
		// Если i-й бит результата равен 1 и расстояние между правой и левой границами диапазона
		// больше, чем 2 в степени i, то бит можно установить в 0, так как он не будет установлен
		// на всех числах в диапазоне
		if (result>>i)&1 == 1 && distance > (1<<i) {
			// Устанавливаем i-й бит результата в 0, используя побитовое XOR
			result = result ^ (1 << i)
		}
	}

	// Возвращаем результат
	return result
}
