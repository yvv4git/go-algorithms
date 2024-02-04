package _07_perfect_number

func checkPerfectNumberV4(num int) bool {
	/*
		METHOD: Bruteforce
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Совершенное число - натуральное число, равное сумме всех своих собственных делителей.
		Собственные делители - это числа, на которые исходное число делится без остатка,
	*/
	return perfectNum(num) == num
}

func perfectNum(perfNum int) int {
	var result = 0

	for i := 1; i < perfNum; i++ {
		if perfNum%i == 0 { // Если число(perfNum) делится на i без остатка, то i является его делителем. Значит можно добавить к результату.
			result = result + i
		}
	}

	return result
}
