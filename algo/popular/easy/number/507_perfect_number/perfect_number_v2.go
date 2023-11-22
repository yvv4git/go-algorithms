package _07_perfect_number

func checkPerfectNumberV2(num int) bool {
	/*
		Time complexity: O(logN), ии, говорит, что O(sqrt(n))
		Space complexity: O(1)

		Time Complexity:
		Time complexity этого алгоритма - O(sqrt(n)), где n - это входное число.
		Это связано с тем, что в цикле for происходит проверка условия left < right, которая выполняется до тех пор,
		пока left не станет больше или равна right. Количество итераций цикла будет равно корню из num,
		так как left увеличивается на 1 на каждой итерации.

		Space Complexity:
		Space complexity этого алгоритма - O(1), потому что алгоритм использует фиксированное количество переменных,
		не зависящих от размера входного числа. Это означает, что алгоритм использует постоянное количество памяти,
		независимо от размера входного числ
	*/
	if num <= 1 { // Если число равно 1, то оно не является совершенным.
		return false
	}

	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 { // Из определения "совершенного числа". Это должны быть делители, сумма которых дает исходное число(num).
			sum = sum + i + (num / i)
		}
	}

	return sum-num == num
}
