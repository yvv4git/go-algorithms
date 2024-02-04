package add_digits

func addDigitsV1(num int) int {
	/*
		Если же число "num" не делится на девять без остатка, то функция использует остаток от деления "num" на девять, как ответ на задачу.
		Это происходит потому, что при делении любого числа на девять с остатком, сумма цифр числа будет равна остатку от деления на девять.
		Например, сумма цифр числа 38 равна 2 (3 + 8 = 11, 1 + 1 = 2, остаток от деления 38 на 9 равен 2).

		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(1)
	*/
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}

	return num % 9
}
