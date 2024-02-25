package main

// Функция myPowV2 вычисляет степень числа x в степени n.
func myPowV2(x float64, n int) float64 {
	// Инициализация переменной pow, которая будет хранить результат вычисления степени.
	pow := float64(1)
	// Инициализация переменной a, которая будет использоваться для умножения.
	a := x
	// Инициализация флага minus, который будет установлен в true, если степень n отрицательная.
	minus := false
	// Проверка, является ли степень n отрицательной. Если это так, то устанавливаем флаг minus в true и меняем знак степени на положительный.
	if n < 0 {
		minus = true
		n = -n
	}
	// Цикл, который будет выполняться, пока степень n больше 0.
	for n > 0 {
		// Проверка, является ли степень n нечетной. Если это так, то умножаем pow на a.
		if (n & 1) != 0 {
			pow = pow * a
		}
		// Умножаем a на себя, чтобы получить следующее значение для умножения.
		a = a * a
		// Сдвигаем степень n вправо на 1 бит, чтобы уменьшить ее на 2.
		n >>= 1
	}
	// Если степень была отрицательной, то делим 1 на pow, чтобы получить результат для отрицательной степени.
	if minus {
		pow = 1 / pow
	}
	// Возвращаем результат вычисления степени.
	return pow
}