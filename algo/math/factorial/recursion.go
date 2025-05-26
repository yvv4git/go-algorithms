package factorial

// FactorialByRecursion вычисляет факториал числа n с использованием рекурсии.
func FactorialByRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1 // Базовый случай(база рекурсии)
	}

	// Рекурсивно вызываем функцию для n-1 и умножаем на n
	return FactorialByRecursion(n-1) * n
}
