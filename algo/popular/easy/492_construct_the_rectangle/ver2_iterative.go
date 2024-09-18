package main

func constructRectangleV2(area int) []int {
	/*
		METHOD: Iterative
		Для решения задачи мы используем следующий подход:
		1. Начинаем с начальных значений длины (l) и ширины (w), где l = area и w = 1.
		2. Итерируемся от 1 до корня квадратного из площади.
		3. Для каждого значения i проверяем, делится ли площадь на i без остатка.
		4. Если делится, то обновляем значения длины и ширины, где l = i и w = area / i.
		5. После завершения цикла проверяем, не нужно ли поменять местами длину и ширину, чтобы длина была больше или равна ширине.
		6. Возвращаем результат в виде массива [длина, ширина].

		TIME COMPLEXITY:
		Временная сложность алгоритма составляет O(√n), где n — площадь прямоугольника.
		Это связано с тем, что мы проходим от 1 до корня квадратного из площади.

		SPACE COMPLEXITY:
		Пространственная сложность алгоритма составляет O(1), так как мы используем фиксированное количество дополнительной памяти
		(переменные l, w и i).
	*/
	l, w := area, 1
	for i := 1; i*i <= area; i++ {
		// Проверяем, делится ли площадь на текущее значение i без остатка
		if area%i == 0 {
			// Если делится, то обновляем значения длины и ширины
			l = i
			w = area / i
		}
	}
	// Проверяем, не нужно ли поменять местами длину и ширину, чтобы длина была больше или равна ширине
	if w > l {
		l, w = w, l
	}
	// Возвращаем результат в виде массива [длина, ширина]
	return []int{l, w}
}

// func main() {
//	// Пример использования
//	area := 37
//	result := constructRectangleV2(area)
//	fmt.Println(result) // Вывод: [37 1]
// }