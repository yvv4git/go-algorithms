package two_numbers_matching

/*
Задачка по циклам на проверку совпадения цифр в двух числах

Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа. Программа получает на вход два числа.

Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
Цифры выводятся в порядке их нахождения в первом числе.

Пример: Ввод: 564 8954 Вывод: 5 4
*/

func twoNumbersMatching(a, b int) []int {
	var x, y int
	var result []int

	j := 10000
	for a > 0 {
		x = a / j
		a = a % j
		y = b
		for y > 0 && x > 0 {
			if y%10 == x {
				// fmt.Print(x, " ")
				result = append(result, x)

			}
			y = y / 10
		}
		j = j / 10
	}

	return result
}
