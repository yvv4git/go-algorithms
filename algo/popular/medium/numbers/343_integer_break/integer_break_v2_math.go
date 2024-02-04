package _43_integer_break

func integerBreakV2(n int) int {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(1)
		Space complexity: O(1)

		Time complexity
		Временная сложность функции integerBreak в худшем случае составляет O(1),
		так как внутри функции нет циклов или рекурсивных вызовов, которые бы могли увеличивать временную сложность в зависимости от входных данных.

		Space complexity
		Пространственная сложность функции integerBreak также составляет O(1),
		так как функция не использует дополнительное пространство, которое бы зависело от размера входных данных.

		Этот код решает задачу "Integer Break" с помощью математики.
		Он проверяет три случая: когда n равно 2 или 3, и когда n делится на 3 с остатком.
		В последнем случае он возвращает 3 в степени (n // 3), если n делится на 3 без остатка,
		3 в степени ((n // 3) - 1) * 4, если n делится на 3 с остатком 1, и 3 в степени (n // 3) * 2, если n делится на 3 с остатком 2.
		Если ни одно из условий не выполняется, то функция возвращает 0.
	*/
	// Если n равно 2, то максимальное произведение равно 1.
	if n == 2 {
		return 1
	}
	// Если n равно 3, то максимальное произведение равно 2.
	if n == 3 {
		return 2
	}
	// Если n делится на 3 без остатка, то максимальное произведение равно 3 в степени (n // 3).
	if n%3 == 0 {
		return pow(3, n/3)
	}
	// Если n делится на 3 с остатком 1, то максимальное произведение равно 3 в степени ((n // 3) - 1) * 4.
	if n%3 == 1 {
		return pow(3, (n/3)-1) * 4
	}
	// Если n делится на 3 с остатком 2, то максимальное произведение равно 3 в степени (n // 3) * 2.
	if n%3 == 2 {
		return pow(3, n/3) * 2
	}
	// Если ни одно из условий не выполняется, то функция возвращает 0.
	return 0
}

// Функция pow возводит число a в степень b.
func pow(a, b int) int {
	/*
		TIME COMPLEXITY: O(b), где b - степень числа.
		Space complexity: O(1)

		Time complexity & Space complexity
		Функция pow имеет временную сложность O(b), где b - степень числа, и пространственную сложность O(1),
		так как она использует фиксированное количество переменных для хранения результата и счетчика.

		В целом, временная и пространственная сложность всех функций в этом коде составляет O(1), что является оптимальным.
	*/
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
