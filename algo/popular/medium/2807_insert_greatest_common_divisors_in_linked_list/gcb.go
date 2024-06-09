package _807_insert_greatest_common_divisors_in_linked_list

// Функция для нахождения НОД двух чисел
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
