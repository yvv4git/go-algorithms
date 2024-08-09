package main

// Функция для нахождения НОД двух чисел
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Основная функция, решающая задачу
func canMeasureWater(x int, y int, target int) bool {
	/*
		METHOD:
		- Используется алгоритм Евклида для нахождения наибольшего общего делителя (НОД) чисел x и y.
		- Проверяются условия:
			1. Если target больше суммы x и y, то сразу возвращается false.
			2. Если target равно 0, то возвращается true.
			3. Проверяется, кратно ли target НОД(x, y). Если да, то возвращается true, иначе false.

		TIME COMPLEXITY:
		- Нахождение НОД с использованием алгоритма Евклида имеет временную сложность O(log(min(x, y))).
		- Проверка условий выполняется за константное время O(1).
		Общая временная сложность: O(log(min(x, y)))

		SPACE COMPLEXITY:
		- Пространственная сложность алгоритма O(1), так как используются только несколько переменных и не требуется дополнительной памяти.
	*/
	// Если target больше суммы x и y, то точно невозможно получить target литров
	if target > x+y {
		return false
	}
	// Если target равно 0, то это всегда возможно
	if target == 0 {
		return true
	}
	// Находим НОД x и y
	gcdXY := gcd(x, y)
	// Проверяем, кратно ли target НОД(x, y)
	return target%gcdXY == 0
}

func main() {
	// Пример использования
	x := 3
	y := 5
	z := 4
	result := canMeasureWater(x, y, z)
	println(result) // true
}
