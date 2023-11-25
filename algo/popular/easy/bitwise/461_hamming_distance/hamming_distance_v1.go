package _61_hamming_distance

func hammingDistanceV1(x int, y int) int {
	/*
		Method: Bitwise
		Time complexity: O(n) - где n, это количество бит в наибольшем из чисел. Внутри цикла выполняется побитовый сдвиг.
		Space complexity: O(1) - используется фиксированного количество переменных, не зависимо от данных.

		1. Сначала выполняется операция XOR (xor := x ^ y).
		Операция XOR сравнивает биты двух чисел и возвращает 1, если биты различны, и 0, если биты одинаковы.

		2. Затем создается переменная distance, которая будет хранить количество различающихся битов.

		3. Далее в цикле (for xor != 0) проверяется, равно ли xor нулю. Если нет, то выполняется следующий код.

		4. В теле цикла выполняется побитовый AND (xor & 1). Если последний бит числа xor равен 1, то результат будет 1, иначе - 0.
		Этот результат добавляется к distance (distance += xor & 1).

		5. Затем xor сдвигается вправо на один бит (xor = xor >> 1).
		Это делается для того, чтобы перейти к следующему биту числа.

		6. Цикл повторяется до тех пор, пока xor не станет равным нулю.

		7. После завершения цикла функция возвращает distance, которая содержит количество различающихся битов между числами x и y
	*/
	xor := x ^ y  // В разрядах, где биты отличаются, будет стоять 1
	distance := 0 // Будет хранить количество различающихся бит

	for xor != 0 {
		distance += xor & 1 // Если последний бит числа равен 1, то результат будет 1, иначе 0. Этот результат добавляется к distance
		xor = xor >> 1      // Сдвигаем наши биты вправо на один разряд
	}

	return distance
}
